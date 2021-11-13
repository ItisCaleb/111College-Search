const fs = require('fs')
var XLSX = require('xlsx');
var workbook = XLSX.readFile('111.xls');
var sheet_name_list = workbook.SheetNames;
sheet_name_list.forEach(function(y) {
    var worksheet = workbook.Sheets[y];
    var headers = {};
    var subheaders = {};
    var data = [];
    var lastcol =""
    for(let z in worksheet) {
        if(z[0] === '!') continue;
        //parse out the column, row, and value
        var col = z.replaceAll(/[0-9]/g, '');
        var row = parseInt(z.replaceAll(/[A-Z]/g,''));

        var value = worksheet[z].v;

        //store header names
        if(row == 1) {
            headers[col] = value;
            continue;
        }
        if(row ==2){
            if(!headers[col]) headers[col] = headers[lastcol]
            else lastcol = col
            subheaders[col] = value;
            continue
        }
        if(!headers[col]) continue
        if(!data[row]) data[row]={};
        if(!subheaders[col]){
            if(typeof value === "string") value = value.replaceAll(" ","")
            data[row][headers[col]] = value;
        }else {
            if(!data[row][headers[col]]){
                data[row][headers[col]] = {}
            }
            data[row][headers[col]][subheaders[col]] = value
        }
    }
    //drop those first two rows which are empty
    data.shift();
    data.shift();
    fs.writeFileSync("./111.json",JSON.stringify(data,null,4))
});
