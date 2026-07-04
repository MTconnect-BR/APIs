import { test, expect } from '@playwright/test';

const PLAYGROUND_URL = 'https://website-ten-gold-16.vercel.app/playground';

test.describe('Playground Page', () => {
  test('page loads with correct title', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    await expect(page).toHaveTitle(/Velo/);
  });

  test('main heading exists', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    const heading = page.locator('h1:has-text("PLAYGROUND")');
    await expect(heading).toBeVisible();
  });

  test('subtitle exists', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    const subtitle = page.locator('text=Teste a API Velo em tempo real');
    await expect(subtitle).toBeVisible();
  });

  test('configuration section exists', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    const configSection = page.locator('h2:has-text("Configuração")');
    await expect(configSection).toBeVisible();
  });

  test('request count slider exists', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    const slider = page.locator('input[type="range"]').first();
    await expect(slider).toBeVisible();
  });

  test('HTTP method buttons exist', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    
    const methods = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH'];
    for (const method of methods) {
      const button = page.locator(`button:has-text("${method}")`);
      await expect(button).toBeVisible();
    }
  });

  test('test buttons exist', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    
    const traditionalBtn = page.locator('button:has-text("Testar Tradicional")');
    const veloBtn = page.locator('button:has-text("Testar Velo API")');
    
    await expect(traditionalBtn).toBeVisible();
    await expect(veloBtn).toBeVisible();
  });

  test('HTTP method buttons are clickable', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    
    const getButton = page.locator('button:has-text("GET")');
    await getButton.click();
    
    const postButton = page.locator('button:has-text("POST")');
    await postButton.click();
    
    const deleteButton = page.locator('button:has-text("DELETE")');
    await deleteButton.click();
  });

  test('navigation links exist', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    
    const docsLink = page.locator('a:has-text("Docs")');
    const testsLink = page.locator('a:has-text("Tests")');
    const compareLink = page.locator('a:has-text("Compare")');
    const githubLink = page.locator('a:has-text("GitHub")');
    
    await expect(docsLink).toBeVisible();
    await expect(testsLink).toBeVisible();
    await expect(compareLink).toBeVisible();
    await expect(githubLink).toBeVisible();
  });

  test('how it works section exists', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    
    const howItWorks = page.locator('h2:has-text("Como funciona")');
    await expect(howItWorks).toBeVisible();
    
    const traditionalApi = page.locator('text=API Tradicional');
    const veloApi = page.locator('text=Velo API');
    
    await expect(traditionalApi).toBeVisible();
    await expect(veloApi).toBeVisible();
  });

  test('footer exists', async ({ page }) => {
    await page.goto(PLAYGROUND_URL);
    
    const copyright = page.locator('text=© 2024 Velo API');
    const license = page.locator('text=MIT License');
    
    await expect(copyright).toBeVisible();
    await expect(license).toBeVisible();
  });
});
