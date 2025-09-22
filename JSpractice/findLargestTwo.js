function findLargestTwo(arrNums) {
  let largest = -Infinity;
  let secondLargest = -Infinity;
  if (secondLargest > largest) {
    secondLargest, (largest = largest), secondLargest;
  }
  for (let i = 0; i < arrNums.length; i++) {
    if (arrNums[i] >= largest) {
      secondLargest = largest;
      largest = arrNums[i];
    } else if (arrNums[i] > secondLargest && secondLargest < largest) {
      secondLargest = arrNums[i];
    }
  }
  return [largest, secondLargest];
}

let arrNums = [
  100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
  1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400, 2500,
];


console.log(findLargestTwo(arrNums));
