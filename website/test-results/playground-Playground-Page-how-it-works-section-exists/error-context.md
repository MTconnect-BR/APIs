# Instructions

- Following Playwright test failed.
- Explain why, be concise, respect Playwright best practices.
- Provide a snippet of code with the fix, if possible.

# Test info

- Name: playground.spec.ts >> Playground Page >> how it works section exists
- Location: tests/playground.spec.ts:82:7

# Error details

```
Error: expect(locator).toBeVisible() failed

Locator: locator('text=Velo API')
Expected: visible
Error: strict mode violation: locator('text=Velo API') resolved to 3 elements:
    1) <button class="px-6 py-3 bg-primary hover:bg-primary/90 disabled:bg-primary/50 text-white rounded-lg font-medium transition flex items-center gap-2">🚀 Testar Velo API</button> aka getByRole('button', { name: '🚀 Testar Velo API' })
    2) <strong class="text-foreground">Velo API:</strong> aka getByText('Velo API:')
    3) <span>© 2024 Velo API</span> aka getByText('© 2024 Velo API')

Call log:
  - Expect "toBeVisible" with timeout 5000ms
  - waiting for locator('text=Velo API')

```

# Page snapshot

```yaml
- generic [active] [ref=e1]:
  - main [ref=e2]:
    - generic [ref=e4]:
      - link "Velo" [ref=e5] [cursor=pointer]:
        - /url: /
      - navigation [ref=e6]:
        - link "Docs" [ref=e7] [cursor=pointer]:
          - /url: /#docs
        - link "Tests" [ref=e8] [cursor=pointer]:
          - /url: /#tests
        - link "Compare" [ref=e9] [cursor=pointer]:
          - /url: /#compare
        - link "Playground" [ref=e10] [cursor=pointer]:
          - /url: /playground
        - link "GitHub" [ref=e11] [cursor=pointer]:
          - /url: https://github.com/MTconnect-BR/APIs
    - generic [ref=e12]:
      - generic [ref=e13]:
        - heading "PLAYGROUND" [level=1] [ref=e14]
        - paragraph [ref=e15]: Teste a API Velo em tempo real
      - generic [ref=e16]:
        - heading "Configuração" [level=2] [ref=e17]
        - generic [ref=e18]:
          - generic [ref=e19]:
            - generic [ref=e20]: "Requests: 50"
            - slider [ref=e21]: "50"
            - generic [ref=e22]:
              - generic [ref=e23]: "10"
              - generic [ref=e24]: "200"
          - generic [ref=e25]:
            - generic [ref=e26]: "Concorrentes: 5"
            - slider [ref=e27]: "5"
            - generic [ref=e28]:
              - generic [ref=e29]: "1"
              - generic [ref=e30]: "50"
        - generic [ref=e31]:
          - generic [ref=e32]: Métodos HTTP
          - generic [ref=e33]:
            - button "GET" [ref=e34]
            - button "POST" [ref=e35]
            - button "PUT" [ref=e36]
            - button "DELETE" [ref=e37]
            - button "PATCH" [ref=e38]
        - generic [ref=e39]:
          - button "⚡ Testar Tradicional" [ref=e40]
          - button "🚀 Testar Velo API" [ref=e41]
      - generic [ref=e42]:
        - heading "Como funciona" [level=2] [ref=e43]
        - generic [ref=e44]:
          - paragraph [ref=e45]:
            - strong [ref=e46]: "API Tradicional:"
            - text: Requests diretos para APIs públicas (jsonplaceholder.typicode.com) sem otimizações de gateway.
          - paragraph [ref=e47]:
            - strong [ref=e48]: "Velo API:"
            - text: Requests passando pelo gateway Velo com rate limiting otimizado, cache de respostas e balanceamento de carga.
          - paragraph [ref=e49]:
            - strong [ref=e50]: "Métricas:"
            - text: Todas as métricas são coletadas em tempo real via
            - code [ref=e51]: performance.now()
            - text: e headers HTTP.
    - generic [ref=e53]:
      - generic [ref=e54]: © 2024 Velo API
      - generic [ref=e55]:
        - link "GitHub" [ref=e56] [cursor=pointer]:
          - /url: https://github.com/MTconnect-BR/APIs
        - generic [ref=e57]: MIT License
  - alert [ref=e58]
```

# Test source

```ts
  1   | import { test, expect } from '@playwright/test';
  2   | 
  3   | const PLAYGROUND_URL = 'https://website-ten-gold-16.vercel.app/playground';
  4   | 
  5   | test.describe('Playground Page', () => {
  6   |   test('page loads with correct title', async ({ page }) => {
  7   |     await page.goto(PLAYGROUND_URL);
  8   |     await expect(page).toHaveTitle(/Velo/);
  9   |   });
  10  | 
  11  |   test('main heading exists', async ({ page }) => {
  12  |     await page.goto(PLAYGROUND_URL);
  13  |     const heading = page.locator('h1:has-text("PLAYGROUND")');
  14  |     await expect(heading).toBeVisible();
  15  |   });
  16  | 
  17  |   test('subtitle exists', async ({ page }) => {
  18  |     await page.goto(PLAYGROUND_URL);
  19  |     const subtitle = page.locator('text=Teste a API Velo em tempo real');
  20  |     await expect(subtitle).toBeVisible();
  21  |   });
  22  | 
  23  |   test('configuration section exists', async ({ page }) => {
  24  |     await page.goto(PLAYGROUND_URL);
  25  |     const configSection = page.locator('h2:has-text("Configuração")');
  26  |     await expect(configSection).toBeVisible();
  27  |   });
  28  | 
  29  |   test('request count slider exists', async ({ page }) => {
  30  |     await page.goto(PLAYGROUND_URL);
  31  |     const slider = page.locator('input[type="range"]').first();
  32  |     await expect(slider).toBeVisible();
  33  |   });
  34  | 
  35  |   test('HTTP method buttons exist', async ({ page }) => {
  36  |     await page.goto(PLAYGROUND_URL);
  37  |     
  38  |     const methods = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH'];
  39  |     for (const method of methods) {
  40  |       const button = page.locator(`button:has-text("${method}")`);
  41  |       await expect(button).toBeVisible();
  42  |     }
  43  |   });
  44  | 
  45  |   test('test buttons exist', async ({ page }) => {
  46  |     await page.goto(PLAYGROUND_URL);
  47  |     
  48  |     const traditionalBtn = page.locator('button:has-text("Testar Tradicional")');
  49  |     const veloBtn = page.locator('button:has-text("Testar Velo API")');
  50  |     
  51  |     await expect(traditionalBtn).toBeVisible();
  52  |     await expect(veloBtn).toBeVisible();
  53  |   });
  54  | 
  55  |   test('HTTP method buttons are clickable', async ({ page }) => {
  56  |     await page.goto(PLAYGROUND_URL);
  57  |     
  58  |     const getButton = page.locator('button:has-text("GET")');
  59  |     await getButton.click();
  60  |     
  61  |     const postButton = page.locator('button:has-text("POST")');
  62  |     await postButton.click();
  63  |     
  64  |     const deleteButton = page.locator('button:has-text("DELETE")');
  65  |     await deleteButton.click();
  66  |   });
  67  | 
  68  |   test('navigation links exist', async ({ page }) => {
  69  |     await page.goto(PLAYGROUND_URL);
  70  |     
  71  |     const docsLink = page.locator('a:has-text("Docs")');
  72  |     const testsLink = page.locator('a:has-text("Tests")');
  73  |     const compareLink = page.locator('a:has-text("Compare")');
  74  |     const githubLink = page.locator('a:has-text("GitHub")');
  75  |     
  76  |     await expect(docsLink).toBeVisible();
  77  |     await expect(testsLink).toBeVisible();
  78  |     await expect(compareLink).toBeVisible();
  79  |     await expect(githubLink).toBeVisible();
  80  |   });
  81  | 
  82  |   test('how it works section exists', async ({ page }) => {
  83  |     await page.goto(PLAYGROUND_URL);
  84  |     
  85  |     const howItWorks = page.locator('h2:has-text("Como funciona")');
  86  |     await expect(howItWorks).toBeVisible();
  87  |     
  88  |     const traditionalApi = page.locator('text=API Tradicional');
  89  |     const veloApi = page.locator('text=Velo API');
  90  |     
  91  |     await expect(traditionalApi).toBeVisible();
> 92  |     await expect(veloApi).toBeVisible();
      |                           ^ Error: expect(locator).toBeVisible() failed
  93  |   });
  94  | 
  95  |   test('footer exists', async ({ page }) => {
  96  |     await page.goto(PLAYGROUND_URL);
  97  |     
  98  |     const copyright = page.locator('text=© 2024 Velo API');
  99  |     const license = page.locator('text=MIT License');
  100 |     
  101 |     await expect(copyright).toBeVisible();
  102 |     await expect(license).toBeVisible();
  103 |   });
  104 | });
  105 | 
```