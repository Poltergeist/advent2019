function simple(mass) {
  return Math.floor(mass / 3) - 2;
}
function advanced(mass) {
  let fuel = simple(mass);

  let additionalFuel = fuel;

  while ((n = simple(additionalFuel)) > 0) {
    fuel += n;
    additionalFuel = n;
  }

  return fuel;
}

module.exports = { simple, advanced };
