import type { GlobalThemeOverrides } from "naive-ui";

const themeOverrides: GlobalThemeOverrides = {
	Form: {
		labelTextColor: "var(--text-color)"
	},
	Input: {
		color: "var(--text-color)",
		colorFocusError: "var(--text-color)",
		borderRadius: "5px",
		border: "0px",
		textColor: "var(--primary-dark-color)",
		placeholderColor: "var(--secondary-color)",
		colorFocus: "var(--text-color)",
		borderHover: "1px solid var(--accent-color)",
		borderFocus: "1px solid var(--accent-color)",
		caretColor: "var(--secondary-color)"
	},
	Select: {
		peers: {
			InternalSelection: {
				color: "var(--text-color)",
				borderRadius: "5px",
				border: "0px",
				textColor: "var(--primary-dark-color)",
				placeholderColor: "var(--secondary-color)",
				borderHover: "1px solid var(--accent-color)",
				borderFocus: "1px solid var(--accent-color)",
				borderActive: "1px solid var(--accent-color)"
			},
			InternalSelectMenu: {
				color: "var(--text-color)",
				borderRadius: "5px",
				optionTextColor: "var(--primary-dark-color)",
				optionTextColorActive: "var(--secondary-color)"
			}
		}
	}
};

export default themeOverrides;
