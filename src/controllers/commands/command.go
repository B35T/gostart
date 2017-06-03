package commands

import (
	"os/exec"
	"gostart/src/models"
	"log"
	"os"
	"io/ioutil"
)

var name = ""

func GetPWD() (string, error){
	cmd := exec.Command("pwd")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(output), err
}

func GoStartMVC(project string) (string, string ,error) {
	mvc := models.MVC{
		Models:"models",
		Views:"views",
		Controllers:"controllers",
		Router:"router",
		Config:"configs",
	}

	file := models.MVC{
		Models:"models/models.go",
		Views:"views/views.go",
		Controllers:"controllers/controller.go",
		Router:"router/router.go",
		Config:"configs/configs.go",
	}

	cmd := exec.Command("mkdir", mvc.Controllers, mvc.Views, mvc.Models, mvc.Config, mvc.Router)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return "","", err
	}

	touch := exec.Command("touch","main.go", file.Models, file.Router, file.Config, file.Views, file.Controllers)
	_ , err = touch.CombinedOutput()
	if err != nil {
		return "", "", err
	}

	move := name+"/src/"+project
	mk := exec.Command("mkdir", project)
	_, err = mk.CombinedOutput()
	if err != nil {
		return "","", err
	}

	mv := exec.Command("mv",project , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}

	mv = exec.Command("mv","models" , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}
	mv = exec.Command("mv","views" , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}
	mv = exec.Command("mv","controllers" , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}
	mv = exec.Command("mv","router" , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}
	mv = exec.Command("mv","configs" , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}

	mv = exec.Command("mv","main.go" , move)
	_,err = mv.CombinedOutput()
	if err != nil {
		return "", "", err
	}

	ls := exec.Command("ls", move)
	output_ls,_ := ls.CombinedOutput()

	return string(output), string(output_ls), err
}

func GoENV() (string, error) {
	println(name)

	pwd := name
	env_test := " export GOPATH=$HOME/"+ pwd +
		"\n export PATH=$PATH:$GOROOT/bin:$GOPATH/bin"
	b := []byte(env_test)

	_, err := os.Create("profile")
	if err != nil {
		return "",err
	}

	err = ioutil.WriteFile("profile", b, 0644)
	if err != nil {
		return "", err
	}

	env := models.ENV{
		PATH:"PATH=$PATH:$GOROOT/bin:$GOPATH/bin",
		GOPATH:"GOPATH=$HOME"+pwd,
	}

	cmd := exec.Command("export", env.PATH)
	cmd = exec.Command("export", env.GOPATH)
	cmd = exec.Command("go","env")
	input, err := cmd.CombinedOutput()
	if err != nil {
		return "",err
	}

	return string(input), err
}

func Create(arg string) (string,error ) {
	name = arg

	mk := exec.Command("mkdir", arg)
	_, err := mk.CombinedOutput()
	if err != nil {
		return "", err
	}

	mk = exec.Command("mkdir", "bin", "src")
	_, err = mk.CombinedOutput()
	if err != nil {
		return "",err
	}
	mv := exec.Command("mv", "bin", arg)
	_, err = mv.CombinedOutput()
	if err != nil {
		return "",err
	}
	mv = exec.Command("mv", "src", arg)
	_, err = mv.CombinedOutput()
	if err != nil {
		return "",err
	}



	return "created", err
}


