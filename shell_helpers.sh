function xe() {
    set -o allexport
    source "$1"
    set +o allexport
}


function ee() {
    eval $(grep -v "^#" $1) "${@:2}"
}


function mkdev() {
    ee .env.dev make "${@:1}"
}
