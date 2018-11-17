const fs = require('fs');

function mkdir(n) {
  for (let i = 1; i <= n; i++) {
    console.log(i);
    if(!fs.existsSync(`exer${i}`)) {
      fs.mkdirSync(`exer${i}`);
    }
  }
}

mkdir(65);
