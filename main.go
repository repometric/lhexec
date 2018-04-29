package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/repometric/lhexec/api"
	"github.com/urfave/cli"
)

// rpcRequest represents a RPC request.
// rpcRequest implements the io.ReadWriteCloser interface.
type rpcRequest struct {
	r    io.Reader     // holds the JSON formated RPC request
	rw   io.ReadWriter // holds the JSON formated RPC response
	done chan bool     // signals then end of the RPC request
}

// NewRPCRequest returns a new rpcRequest.
func NewRPCRequest(r io.Reader) *rpcRequest {
	var buf bytes.Buffer
	done := make(chan bool)
	return &rpcRequest{r, &buf, done}
}

// Read implements the io.ReadWriteCloser Read method.
func (r *rpcRequest) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

// Write implements the io.ReadWriteCloser Write method.
func (r *rpcRequest) Write(p []byte) (n int, err error) {
	return r.rw.Write(p)
}

// Close implements the io.ReadWriteCloser Close method.
func (r *rpcRequest) Close() error {
	r.done <- true
	return nil
}

// Call invokes the RPC request, waits for it to complete, and returns the results.
func (r *rpcRequest) Call() io.Reader {
	go jsonrpc.ServeConn(r)
	<-r.done
	return r.rw
}

const appVersion = "0.0.2"
const apiPort = ":1234"

func main() {
	app := cli.NewApp()
	app.Version = appVersion
	app.Usage = "Linterhub Engine Executor"
	app.Commands = []cli.Command{
		/*{
			Name:    "analyze",
			Aliases: []string{"a"},
			Usage:   "Runs code analysis",
			Action: func(c *cli.Context) error {
				var context = analyze.Context{
					Project:     c.String("project"),
					File:        c.String("file"),
					Folder:      c.String("folder"),
					Environment: c.String("environment"),
					Engine:      c.StringSlice("engine"),
					Stdin:       c.Bool("stdin"),
				}
				if len(context.Engine) == 0 || len(context.Project) == 0 {
					cli.ShowCommandHelp(c, "analyze")
				} else {
					var res []byte
					if context.Stdin {
						reader := bufio.NewReader(os.Stdin)
						var output []rune

						for {
							input, _, err := reader.ReadRune()
							if err != nil && err == io.EOF {
								break
							}
							output = append(output, input)
						}

						context.StdinContent = string(output)
					}

					analyzeResult, _ = analyze.Run(context)
					res, _ = json.MarshalIndent(analyzeResult, "", "    ")
					fmt.Println(string(res))
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "engine,e",
					Usage: "List of the engines to be analyzed",
				},
				cli.StringFlag{
					Name:  "project,p",
					Usage: "Path to the project",
				},
				cli.StringFlag{
					Name:  "file,f",
					Usage: "Specific file for analysis",
				},
				cli.StringFlag{
					Name:  "folder,F",
					Usage: "Relative path to the specific folder",
				},
				cli.StringFlag{
					Name:  "environment,env",
					Usage: "The way how to analyze (local by default)",
				},
				cli.BoolFlag{
					Name:  "stdin",
					Usage: "This argument says that standard input will be analyzed (false by default)",
				},
			},
		},*/
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Runs jsonrpc service",
			Action: func(c *cli.Context) error {
				api := api.GetInstance()
				rpc.Register(api)
				http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
					defer req.Body.Close()
					w.Header().Set("Content-Type", "application/json")
					res := NewRPCRequest(req.Body).Call()
					io.Copy(w, res)
				})
				log.Fatal(http.ListenAndServe(apiPort, nil))
				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Returns current component version",
			Action: func(c *cli.Context) error {
				fmt.Println(c.App.Version)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
