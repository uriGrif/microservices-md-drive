import type { GlobalThemeOverrides } from "naive-ui";

const themeOverrides: GlobalThemeOverrides = {
    Form: {
        labelTextColor: "var(--text-color)"
    },
    Input: {
        color: "var(--text-color)",
        colorFocusError: "var(--text-color)",
        borderRadius: "var(--border-radius)",
        textColor: "var(--primary-dark-color)",
        placeholderColor: "var(--primary-dark-color)",
        colorFocus: "var(--text-color)",
        borderHover: "1px solid var(--accent-color)",
        borderFocus: "1px solid var(--accent-color)",
        caretColor: "var(--secondary-color)"
    },
};

export default themeOverrides;
