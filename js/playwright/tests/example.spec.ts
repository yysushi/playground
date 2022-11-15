import { expect, test } from "@playwright/test";

test("homepage has title and links to intro page", async ({ page }) => {
  await page.goto("https://playwright.dev/");

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/Playwright/);

  // create a locator
  const getStarted = page.getByRole("link", { name: "Get started" });

  // Expect an attribute "to be strictly equal" to the value.
  await expect(getStarted).toHaveAttribute("href", "/docs/intro");

  // Click the get started link.
  await getStarted.click();

  // Expects the URL to contain intro.
  await expect(page).toHaveURL(/.*intro/);
});

test("hello", async ({ page }) => {
  await page.goto("https://google.com");

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/oog/);
});

// http://localhost:10000/JapanEast1-A/vnc_auto.html?path=JapanEast1-A
test("hello2", async ({ page }) => {
  await page.goto(
    "http://localhost:10000/JapanEast1-A/vnc_auto.html?path=JapanEast1-A",
  );

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/noVNC/);

        sagss;
});
