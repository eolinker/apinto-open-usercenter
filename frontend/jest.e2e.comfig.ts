module.exports = {
    roots: ['<rootDir>/src'],
    preset: 'jest-playwright-preset',
    transform: {
      '^.+\\.ts$': 'ts-jest',
      '^.+\\.js$': 'babel-jest'
  
    },
    testPathIgnorePatterns: ['/node_modules/', 'dist', 'lib'],
    testMatch: ['**/**/*.test.{js,jsx,ts,tsx}'],
    verbose: true,
    setupFilesAfterEnv: ['<rootDir>/test/setup-e2e.ts'],
    moduleNameMapper: {
      '~src/(.*)$':'<rootDir>/src/$1',
      // angular 13+
      //'@angular/core/testing': '<rootDir>/node_modules/@angular/core/fesm2015/testing.mjs',
      //'@angular/platform-browser-dynamic/testing': '<rootDir>/node_modules/@angular/platform-browser-dynamic/fesm2015/testing.mjs',
      //'@angular/platform-browser/testing': '<rootDir>/node_modules/@angular/platform-browser/fesm2015/testing.mjs',
    },
  }