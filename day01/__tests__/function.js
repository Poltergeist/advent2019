const simple = require("../function").simple;
const advanced = require("../function").advanced;

describe("day 1", () => {
  describe("simple", () => {
    [[12, 2], [14, 2], [1969, 654], [100756, 33583]].forEach(([mass, result]) =>
      it(`mass of ${mass} resolves to ${result}`, () =>
        expect(simple(mass)).toEqual(result))
    );
  });
  describe("advanced", () => {
    [[12, 2], [14, 2], [1969, 966], [100756, 50346]].forEach(([mass, result]) =>
      it(`mass of ${mass} resolves to ${result}`, () =>
        expect(advanced(mass)).toEqual(result))
    );
  });
});
