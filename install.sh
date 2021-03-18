BIN="_out/pbld"

if command -v go > /dev/null 2>&1 ; then
   echo "found $(go version)"
else
   echo "go not found. Please install golang using your systems package manager"
fi

if [[ -e ${BIN}  ]] ; then
   cp _out/pbld /usr/local/bin
fi
