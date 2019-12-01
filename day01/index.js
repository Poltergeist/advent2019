const fs = require("fs");
const simple = require("./function.js").advanced;

fs.readFile(`${__dirname}/input`, "utf-8", (error, data) => {
  console.log(
    data.split("\n").reduce((acc, cur) => {
      return acc + simple(parseInt(cur));
    }, 0)
  );
});
