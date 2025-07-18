import { devices, type PlaywrightTestConfig } from "@playwright/test";
import { ADMIN_STORAGE_STATE } from "./tests/constants";

const config: PlaywrightTestConfig = {
  webServer: {
    command: "npm run build && npm run preview",
    port: 3000,
    reuseExistingServer: !process.env.CI,
    timeout: 120_000,
    cwd: "../web-admin",
  },
  retries: 0,
  reporter: "html",
  use: {
    baseURL: "http://localhost:3000",
    ...devices["Desktop Chrome"],
    trace: "retain-on-failure",
    video: "retain-on-failure",
  },
  testDir: "tests",
  projects: [
    {
      name: "setup",
      testMatch: "setup.ts",
      ...(process.env.E2E_NO_TEARDOWN ? undefined : { teardown: "teardown" }),
    },
    {
      name: "teardown",
      testMatch: "teardown.ts",
      use: {
        storageState: ADMIN_STORAGE_STATE,
      },
    },
    {
      name: "e2e",
      dependencies: process.env.E2E_NO_SETUP_OR_TEARDOWN ? [] : ["setup"],
      testIgnore: "/setup",
      use: {
        storageState: ADMIN_STORAGE_STATE,
      },
    },
  ],
};

export default config;
