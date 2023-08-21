#! /bin/bash

if [ $# -ne 1 ];
    then echo "illegal number of parameters"
else
    echo "Updating views using $1-frontend"

    cd "$1-frontend/"

    npm run build

    case $1 in 
        vue)
            cp -r dist/* ../views/
            ;;

        *)
            cp -r build/* ../views/
            ;;
    esac

fi