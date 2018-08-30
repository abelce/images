package commandHandler

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"os"
	"time"
	"strings"
	"strconv"
	"log"
	"path/filepath"

	"images/application/command"
	"images/domain/model"

	"github.com/fogleman/primitive/primitive"
	"github.com/nfnt/resize"
	"github.com/satori/go.uuid"
)

var (
	Input      string
	Outputs    flagArray
	Background string
	Configs    shapeConfigArray
	Alpha      int
	InputSize  int
	OutputSize int
	Mode       int
	Workers    int
	Nth        int
	Repeat     int
	V, VV      bool
)

type flagArray []string

func (i *flagArray) String() string {
	return strings.Join(*i, ", ")
}

func (i *flagArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type shapeConfig struct {
	Count  int
	Mode   int
	Alpha  int
	Repeat int
}

type shapeConfigArray []shapeConfig

func (i *shapeConfigArray) String() string {
	return ""
}

func (i *shapeConfigArray) Set(value string) error {
	n, _ := strconv.ParseInt(value, 0, 0)
	*i = append(*i, shapeConfig{int(n), Mode, Alpha, Repeat})
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateSvg(Input string) (svgurl string) {

	start := time.Now()
	frame := 0
	Nth = 1
	OutputSize = 512
	Workers = 0
	Mode = 1
	Alpha = 128
	Repeat = 0
	Configs = append(Configs, shapeConfig{100, Mode, Alpha, Repeat});
	fid, _ := uuid.NewV4()
	output := "/data/upload_files/" + fid.String() + ".svg"

	if len(Configs) == 1 {
		Configs[0].Mode = Mode
		Configs[0].Alpha = Alpha
		Configs[0].Repeat = Repeat
	}

	input, err := primitive.LoadImage(Input)
	check(err)
	size := uint(InputSize)
	if size > 0 {
		input = resize.Thumbnail(size, size, input, resize.Bilinear)
	}
	var bg = primitive.MakeColor(primitive.AverageImageColor(input))
	newModel := primitive.NewModel(input, bg, OutputSize, Workers)
	primitive.Log(1, "%d: t=%.3f, score=%.6f\n", 0, 0.0, newModel.Score)


	for j, config := range Configs {
		primitive.Log(1, "count=%d, mode=%d, alpha=%d, repeat=%d\n",
			config.Count, config.Mode, config.Alpha, config.Repeat)

		for i := 0; i < config.Count; i++ {
			frame++

			// find optimal shape and add it to the newModel
			t := time.Now()
			n := newModel.Step(primitive.ShapeType(config.Mode), config.Alpha, config.Repeat)
			nps := primitive.NumberString(float64(n) / time.Since(t).Seconds())
			elapsed := time.Since(start).Seconds()
			primitive.Log(1, "%d: t=%.3f, score=%.6f, n=%d, n/s=%s\n", frame, elapsed, newModel.Score, n, nps)

			// write output image(s)
			// for _, output := range Outputs {
				ext := strings.ToLower(filepath.Ext(output))
				percent := strings.Contains(output, "%")
				saveFrames := percent && ext != ".gif"
				saveFrames = saveFrames && frame%Nth == 0
				last := j == len(Configs)-1 && i == config.Count-1
				if saveFrames || last {
					path := output
					if percent {
						path = fmt.Sprintf(output, frame)
					}
					primitive.Log(1, "writing %s\n", path)
					switch ext {
					default:
						check(fmt.Errorf("unrecognized file extension: %s", ext))
					case ".svg":
						check(primitive.SaveFile(path, newModel.SVG()))
					}
				}
			// }
		}
	}

	return output
}

type CreateImage struct {
	ImageRepository model.Repository
	QueryService    model.QueryService
}

// (title, markdowncontent, private, tags, status, categories, typ, description string)
func (h CreateImage) Handle(c command.CreateImage) (*model.Image, error) {

	// 设置图片的宽度，高度
	u, err := url.Parse(c.Url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	path := "/data/upload_files" + u.Path
	file, _ := os.Open(path)
	defer file.Close()
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.Width = img.Width
	c.Height = img.Height
	c.SvgUrl = CreateSvg(c.Url)
	image := model.NewImage(c.Url, c.Width, c.Height, c.SvgUrl)
	err = h.ImageRepository.Save(image)
	if err != nil {
		return nil, err
	}
	return image, nil
}
