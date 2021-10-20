import { defineConfig } from "vite-plugin-windicss";
import typography from "windicss/plugin/typography";

function cssVarRgbHelper(cssVariable: string) {
  return ({
    opacityVariable,
    opacityValue,
  }: {
    opacityVariable: string;
    opacityValue: number;
  }) => {
    if (opacityValue !== undefined) {
      return `rgba(var(--${cssVariable}), ${opacityValue})`;
    }
    if (opacityVariable !== undefined) {
      return `rgba(var(--${cssVariable}), var(${opacityVariable}, 1))`;
    }
    return `rgb(var(--${cssVariable}))`;
  };
}

export default defineConfig({
  darkMode: "class",
  plugins: [typography()],
  theme: {
    extend: {
      colors: {
        primary: {
          100: cssVarRgbHelper("primary-100"),
          200: cssVarRgbHelper("primary-200"),
          300: cssVarRgbHelper("primary-300"),
          400: cssVarRgbHelper("primary-400"),
          500: cssVarRgbHelper("primary-500"),
          600: cssVarRgbHelper("primary-600"),
          700: cssVarRgbHelper("primary-700"),
          800: cssVarRgbHelper("primary-800"),
          900: cssVarRgbHelper("primary-900"),
        },
        secondary: {
          100: cssVarRgbHelper("secondary-100"),
          200: cssVarRgbHelper("secondary-200"),
          300: cssVarRgbHelper("secondary-300"),
          400: cssVarRgbHelper("secondary-400"),
          500: cssVarRgbHelper("secondary-500"),
          600: cssVarRgbHelper("secondary-600"),
          700: cssVarRgbHelper("secondary-700"),
          800: cssVarRgbHelper("secondary-800"),
          900: cssVarRgbHelper("secondary-900"),
        },
        surface: {
          100: cssVarRgbHelper("surface-100"),
          200: cssVarRgbHelper("surface-200"),
          300: cssVarRgbHelper("surface-300"),
          400: cssVarRgbHelper("surface-400"),
          500: cssVarRgbHelper("surface-500"),
          600: cssVarRgbHelper("surface-600"),
          700: cssVarRgbHelper("surface-700"),
          800: cssVarRgbHelper("surface-800"),
          900: cssVarRgbHelper("surface-900"),
        },
      },
    },
  },
});
