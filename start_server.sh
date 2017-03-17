if [[ ! -e out/ ]];
		then
			mkdir out/
		fi 

		go build -o out/build app/workflow/main.go; ./out/build
