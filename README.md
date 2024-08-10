# cipherCode-cli

Installation:

Si cela n'est pas encore fait ajouter les exports suivants dans votre fichier ~/.zshrc:
 export GOPATH=$HOME/go
 export GOROOT=/usr/local/go ** remplacer user par votre nom d'utilisateur**
 export GOBIN=$GOPATH/bin
 export PATH=$PATH:$GOPATH
 export PATH=$PATH:$GOROOT/bin
 
1)Cloner le repo
 
2)Une fois dans votre dossier comprenant le repo tapez:
go mod tidy
go build
go install


**si votre mac bloque l'ewecution mais normalement c'est pas le cas tapez :**
chmod +x cipherCode

3)Tapez dans votre terminal cipherCode et laissez vous guider !
