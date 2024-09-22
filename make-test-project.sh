mkdir -p test-project/
cp main.go test-project/
sed -i 's/package measuregc/package main/g' test-project/main.go
echo -e '\nfunc main() {\n\tStart()\n\tselect {}\n}' >> test-project/main.go
echo "test project created in ./test-project"
