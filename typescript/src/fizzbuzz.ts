export function initializeFizzBuzz(
  list: HTMLOListElement,
  button: HTMLButtonElement
): void {
  console.debug("initializeFizzBuzz", list, button);

  let counter = create_counter();

  button.addEventListener("click", () => {
    const item = counter.next();
    // There's no way to express a type of infinite generator in Typescript
    // and there's a reason: https://stackoverflow.com/questions/72167179/how-to-specify-type-of-an-infinite-generator
    if (item.done) throw Error("unreachable");
    const text = fizzbuzz(item.value);
    const li = document.createElement("li");
    // innerText works in browser but not in test (jsdom) https://github.com/jsdom/jsdom/issues/1245
    li.textContent = text;
    list.appendChild(li);
  });
}

function* create_counter() {
  let i = 1;
  while (true) {
    yield i;
    i += 1;
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
