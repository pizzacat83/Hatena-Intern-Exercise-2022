export function initializeFizzBuzz(
  list: HTMLOListElement,
  button: HTMLButtonElement
): void {
  console.debug("initializeFizzBuzz", list, button);

  let counter = new Counter();

  button.addEventListener("click", () => {
    const text = fizzbuzz(counter.next());
    const li = document.createElement("li");
    // innerText works in browser but not in test (jsdom) https://github.com/jsdom/jsdom/issues/1245
    li.textContent = text;
    list.appendChild(li);
  });
}

class Counter {
  #next = 1;

  next(): number {
    return this.#next++;
  }
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
