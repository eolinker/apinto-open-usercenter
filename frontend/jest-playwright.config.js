module.exports = {
  launchOptions: {
    headless: !process.argv.includes('--headless=false'),
    viewport: { width: 1440, height: 740 }
  },
  serverOptions: {
    command: 'npm run start',
    port: 4200,
    launchTimeout: 1000000,
    debug: true,
    usedPortAction: 'ignore'
  },
  browsers: ['chromium'],
  devices: [],
  collectCoverage: true,
  config: {
    use: {
      viewport: { width: 1440, height: 720 }
    }
  }
}
