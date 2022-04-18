import filelist from "k6/x/listfile"

export default function() {
let files = filelist.getFiles('/Users/ilyakobelev/Documents')
console.log(files)
}