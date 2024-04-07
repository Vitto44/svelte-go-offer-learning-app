/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		extend: {
			colors: {
				brand: "#004080",
				secondary: {
					100: "#ffb633",
					200: "#FFA400",
				},
				text: {
					primary: "#ffffff",
					gray: "#666666",
				},
			},
			fontFamily: {
				sans: ["Roboto", "Helvetica", "Arial", "sans-serif"],
			},
			boxShadow: {
				primary: "0px 3px 8px rgba(0, 0, 0, 0.24)",
			},
		},
	},
	plugins: [],
};
