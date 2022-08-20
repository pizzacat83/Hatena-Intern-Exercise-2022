export function initializeFizzBuzz(
  list: HTMLOListElement,
  button: HTMLButtonElement
): void {
  console.debug("initializeFizzBuzz", list, button);

  let i = 1;

  button.addEventListener("click", () => {
    const text = fizzbuzz(i);
    const li = document.createElement("li");
    // innerText works in browser but not in test (jsdom) https://github.com/jsdom/jsdom/issues/1245
    li.textContent = text;
    list.appendChild(li);
    i += 1;
  });
}

const fizzbuzz = (i: number): string => {
  if (i % 3 == 0 && i % 5 == 0) {
    return "FizzBuzz";
  } else if (i % 3 == 0) {
    return "Fizz";
  } else if (i % 5 == 0) {
    return "Buzz";
  } else {
    return i.toString();
  }
}
