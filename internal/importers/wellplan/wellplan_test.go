package wellplan

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"html"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// buildDocx creates a minimal .docx whose body holds the given paragraphs and tables.
// A table is a [][]string; a string is a paragraph.
func buildDocx(t *testing.T, parts ...any) []byte {
	t.Helper()
	var body strings.Builder
	for _, p := range parts {
		switch v := p.(type) {
		case string:
			body.WriteString(`<w:p><w:r><w:t>` + html.EscapeString(v) + `</w:t></w:r></w:p>`)
		case [][]string:
			body.WriteString(`<w:tbl>`)
			for _, row := range v {
				body.WriteString(`<w:tr>`)
				for _, cell := range row {
					body.WriteString(`<w:tc><w:p><w:r><w:t>` + html.EscapeString(cell) + `</w:t></w:r></w:p></w:tc>`)
				}
				body.WriteString(`</w:tr>`)
			}
			body.WriteString(`</w:tbl>`)
		}
	}
	doc := `<?xml version="1.0"?><w:document xmlns:w="` + wordNS + `"><w:body>` + body.String() + `</w:body></w:document>`
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	w, err := zw.Create("word/document.xml")
	if err != nil {
		t.Fatal(err)
	}
	w.Write([]byte(doc))
	zw.Close()
	return buf.Bytes()
}

func russianReport(t *testing.T, withSurvey bool) []byte {
	parts := []any{
		"Основная информация по кейсу",
		[][]string{
			{"Компания", "Тест-Бурение"},
			{"Месторождение", "Тестовое", "Кустовая площадка", "12 ГС"},
			{"Скважина", "12", "Ствол", "12"},
			{"Траектория", "Design #1", "Кейс", "2159"},
			{"Забой по стволу", "1,020.00 m", "Забой по вертикали", "990.50 m"},
			{"Начало отсчета", "WELL (copy) @  -19m", "Тип скважины", "Platform"},
		},
		[][]string{{"Жидкость", "Полимер", "Тип", "Mud"}, {"Основа", "Water", "Жидкость основы", "Water"}},
		[][]string{{"Температура(°C)", "Давление(atm)", "Плотность(kg/m³)", "Учет", "PV", "YP", "Fann"}, {"", "", "", "", "", "", "rpm"}, {"70", "1", "1230", "Yes", "20", "22", ""}},
		[][]string{{"Тип секции", "Глубина секции(m)", "Длина секции(m)", "Глубина башмака(m)", "Внутренний диаметр(mm)", "Зазор (mm)", "Эффективный диаметр(mm)", "Коэф. трения", "Линейная вместимость (L/m)", "Объемное Уширение(%)"},
			{"Casing", "500.00", "500.000", "500.00", "226.70", "222.63", "", "0.25", "40.36", ""},
			{"Open Hole", "1020.00", "520.000", "", "215.90", "222.25", "215.90", "", "36.61", "0.00"}},
		// Header table followed by a continuation table whose first row is empty.
		[][]string{{"Тип", "Длина(m)", "Глубина(m)", "Тело элемента", "Стабилизатор / Замок трубы", "Вес(kg/m)", "Материал", "Марка стали", "Класс износа"},
			{"", "", "", "Наруж. диаметр(mm)", "Внутр. Диаметр(mm)", "Средняя длина трубы(m)", "Длина(m)", "Наруж. диаметр(mm)", "Внутр. диаметр(mm)", "", "", "", ""}},
		[][]string{{"", "", "", "", "", "", "", "", "", "", "", "", ""},
			{"Drill Pipe", "1,000", "1,000", "127", "108.61", "9.14", "0.433", "152.4", "82.55", "32.62", "CS_API 5D/7", "X", "2"},
			{"Bit", "0.3", "1,000.3", "215.9", "", "0.3", "", "", "", "100", "", "", ""}},
		[][]string{{"Марка стали", "Минимальный предел текучести (psi)"}, {"X", "105,000"}},
		[][]string{{"Тип секции", "Коэффициенты трения"}, {"Casing", "0.25"}, {"Open Hole", "0.30"}},
		[][]string{{"Глубина долота по стволу", "1000.30 m", "Использовать", "Yes"}, {"Вес блока", "17.00 tonne", "Модель", "No"}},
		[][]string{{"Операции при бурении", "Нагрузка на долото/Затяжка(tonne)", "Момент на долоте(kN-m)"}, {"Вращение на забое", "6.00", "4.1670"}, {"Операции при СПО", "Скорость(m/min)", "Обороты(rpm)"}, {"Спуск", "10.00", "0"}},
		[][]string{{"Условие нагружения", "Отказ по напряжению", "…"}, {"", "Усталость", "…"},
			{"Спуск", "", "", "", "", "", "", "", "0.0000", "0.0", "0.0", "67.30", "1.34", "1956.81", "339.19", "2296.00", "0.00"},
			{"Вращение на забое", "", "", "", "X", "", "", "", "10.1015", "3.2", "1.7", "67.81", "1.31", "1853.51", "442.49", "1886.78", "409.22"}},
	}
	if withSurvey {
		parts = append(parts, [][]string{
			{"Глубина по стволу(m)", "Зенитный угол(°)", "Азимутальный угол(°)", "Глубина по вертикали(m)", "Пространст. интенсивн. искривления(°/30m)", "Абсолют. извил-сть(°/30m)", "Относит. извил-сть(°/30m)", "Смещ.-е(m)", "NS(m)", "EW(m)"},
			{"0.00", "0.00", "1.81", "0.00", "0.000", "0", "0", "0.00", "0.00", "0.00"},
			{"1020.00", "30.00", "10.82", "990.50", "5.510", "0", "0", "100.00", "98.00", "20.00"},
		})
	}
	return buildDocx(t, parts...)
}

func f(p *float64) float64 {
	if p == nil {
		return -1
	}
	return *p
}

func TestParseRussianReport(t *testing.T) {
	report, err := ParseReport("case.docx", russianReport(t, true))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Combine(report, nil)
	if err != nil {
		t.Fatal(err)
	}
	c := r.Case
	if c.Company != "Тест-Бурение" || c.Field != "Тестовое" || c.Site != "12 ГС" || c.Well != "12" || c.Case != "2159" || f(c.MD) != 1020 || f(c.DatumElevation) != -19 {
		t.Fatalf("case info: %+v", c)
	}
	if len(r.String) != 2 || r.String[0].Length != 1000 || f(r.String[0].JointOD) != 152.4 || f(r.String[0].MinYield) != 105000 || r.String[1].BodyID != nil {
		t.Fatalf("string: %+v", r.String)
	}
	if len(r.Holes) != 2 || f(r.Holes[1].FrictionFactor) != 0.30 || f(r.Holes[0].ShoeDepth) != 500 {
		t.Fatalf("holes (friction must fall back to the factor table): %+v", r.Holes)
	}
	if r.Fluid == nil || f(r.Fluid.Density) != 1230 || f(r.Fluid.YieldPoint) != 22 || r.Fluid.Name != "Полимер" {
		t.Fatalf("fluid: %+v", r.Fluid)
	}
	if len(r.Survey) != 2 || r.Survey[1].Inc != 30 || f(r.Survey[1].NS) != 98 || f(r.Survey[1].SubSea) != 990.5+19 {
		t.Fatalf("survey (sub-sea derived from datum): %+v", r.Survey)
	}
	td := r.TorqueDrag
	if td == nil || f(td.BitDepth) != 1000.3 || td.WOB[OpRotatingOnBottom] != 6 || td.TripSpeed[OpTrippingIn] != 10 || len(td.Loads) != 2 {
		t.Fatalf("torque & drag: %+v", td)
	}
	rob := td.Loads[1]
	if rob.Operation != OpRotatingOnBottom || f(rob.SurfaceTorque) != 10.1015 || f(rob.HookLoad) != 67.81 || f(rob.NeutralFromBit) != 409.22 || len(rob.Failures) != 1 {
		t.Fatalf("load row: %+v", rob)
	}
}

func TestReportWithoutSurveyNeedsSurveyFile(t *testing.T) {
	report, err := ParseReport("case.docx", russianReport(t, false))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Combine(report, nil); !errors.Is(err, ErrNoSurvey) {
		t.Fatalf("want ErrNoSurvey, got %v", err)
	}
	survey, err := ParseSurvey("plan.txt", []byte(sampleSurvey))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Combine(report, survey)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Survey) != 3 || len(r.Source) != 2 || r.Case.Company != "Тест-Бурение" || f(r.Survey[0].GlobalN) != 5291656.095 {
		t.Fatalf("combined: %+v", r)
	}
}

const sampleSurvey = "Customer\t04. Тест\t Creation Date 1/1/1990\r\n" +
	"Field\t????\tYour Ref\t\r\n" +
	"Wellhead\t2703\tKelly Bushing Elev.\t104.94\r\n\r\n" +
	"MD\tIncl.\tAzim.\tSub-Sea\tTVD\tLocal N Coord\tLocal E Coord\tGlobal N Coord\tGlobal E Coord\tDogleg\tVertical Section\r\n" +
	"(m)\tDeg.\tDeg.\t(m)\t(m)\t(m)\t(m)\t(m)\t(m)\t(dega/30m)\t(m)\t\r\n\r\n" +
	"160.000   \t0.100     \t337.540   \t55.060    \t160.000    \t0.116      \t-0.119     \t5291656.111  \t10286162.800 \t0.000  \t0.09     \r\n" +
	"150.000   \t0.100     \t337.540   \t45.060    \t150.000    \t0.100      \t-0.113     \t5291656.095  \t10286162.807 \t0.000  \t0.09     \r\n" +
	"350.000   \t2.224     \t236.371   \t245.057   \t349.997    \t0.307      \t-0.404     \t5291656.302  \t10286162.515 \t6.736  \t0.32     \r\n"

func TestParseSurveyText(t *testing.T) {
	s, err := ParseSurvey("plan.txt", []byte(sampleSurvey))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Combine(nil, s)
	if err != nil {
		t.Fatal(err)
	}
	if r.Survey[0].MD != 150 || f(r.Survey[2].DLS) != 6.736 || f(r.Survey[0].SubSea) != 45.06 || f(r.SurveyInfo.KellyBushingElev) != 104.94 {
		t.Fatalf("survey: %+v", r.Survey)
	}
	if r.Case.Field != "Imported field" {
		t.Fatalf("names lost to encoding (????) must fall back: %q", r.Case.Field)
	}
	if !strings.Contains(strings.Join(r.Warnings, "\n"), "sorted") {
		t.Fatalf("out-of-order stations must be reported: %v", r.Warnings)
	}
}

func TestRejectsFeetAndGarbage(t *testing.T) {
	feet := strings.Replace(sampleSurvey, "(m)\tDeg.", "(ft)\tDeg.", 1)
	if _, err := ParseSurvey("ft.txt", []byte(feet)); !errors.Is(err, ErrInvalidFile) {
		t.Fatalf("feet survey must be rejected, got %v", err)
	}
	if _, err := ParseReport("x.docx", []byte("not a zip")); !errors.Is(err, ErrInvalidFile) {
		t.Fatalf("garbage must be rejected, got %v", err)
	}
	if _, err := ParseReport("empty.docx", buildDocx(t, "hello")); !errors.Is(err, ErrInvalidFile) {
		t.Fatalf("a docx that is not a WellPlan report must be rejected, got %v", err)
	}
}

func TestParseNumber(t *testing.T) {
	for in, want := range map[string]float64{"1,854": 1854, "55,000": 55000, "1,984.23": 1984.23, "0,25": 0.25, "-24": -24, "30.00 m": 30, "6.8046 atm": 6.8046} {
		if got, ok := parseNumber(in); !ok || got != want {
			t.Errorf("parseNumber(%q) = %v, %v; want %v", in, got, ok, want)
		}
	}
	for _, in := range []string{"", "NA", "abc"} {
		if _, ok := parseNumber(in); ok {
			t.Errorf("parseNumber(%q) should fail", in)
		}
	}
}

// TestRealReports runs over a local folder of real reports when WELLPLAN_REPORTS_DIR is set.
// Real reports contain client data and are never committed.
func TestRealReports(t *testing.T) {
	dir := os.Getenv("WELLPLAN_REPORTS_DIR")
	if dir == "" {
		t.Skip("set WELLPLAN_REPORTS_DIR to parse real reports")
	}
	files, _ := filepath.Glob(filepath.Join(dir, "*.docx"))
	var complete, noSurvey int
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		report, err := ParseReport(filepath.Base(path), data)
		if err != nil {
			t.Errorf("%s: %v", filepath.Base(path), err)
			continue
		}
		r, err := Combine(report, nil)
		if errors.Is(err, ErrNoSurvey) {
			noSurvey++
			continue
		}
		if err != nil {
			t.Errorf("%s: %v", filepath.Base(path), err)
			continue
		}
		if dump := os.Getenv("WELLPLAN_DUMP"); dump != "" && strings.Contains(path, dump) {
			out, _ := json.MarshalIndent(r, "", " ")
			t.Logf("%s", out)
		}
		loads := 0
		if r.TorqueDrag != nil {
			loads = len(r.TorqueDrag.Loads)
		}
		if len(r.String) == 0 || len(r.Holes) == 0 || loads == 0 {
			t.Logf("partial %s: string=%d holes=%d loads=%d", filepath.Base(path), len(r.String), len(r.Holes), loads)
		} else {
			complete++
		}
	}
	t.Logf("%d files: %d complete, %d need a survey file", len(files), complete, noSurvey)
}
