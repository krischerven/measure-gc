mkdir -p testproj/
cp main.go testproj/
sed -i 's/package measuregc/package main/g' testproj/main.go
echo -e '\nfunc main() {\n\tStart()\n\tselect {}\n}' >> testproj/main.go
