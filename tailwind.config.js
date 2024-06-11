/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
	safelist: [],
	plugins: [require("daisyui")],
	daisyui: {
		themes: ["dark", "sunset", "luxury",
			{
				cool: {
					"primary": "#ff00e4",
					"secondary": "#00e46a",
					"accent": "#009600",
					"neutral": "#0a0207",
					"base-100": "#2a282d",
					"info": "#00feff",
					"success": "#00da70",
					"warning": "#ee9600",
					"error": "#ff6394",
				}
			}
		]
	}
}