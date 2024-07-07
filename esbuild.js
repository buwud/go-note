const esbuild = require("esbuild");
(async () => {
    let ctx = await esbuild.context({
        entryPoints: ["frontend/Application.tsx", "frontend/style.css"],
        outdir: "public/assets",
        bundle: true,
    });
    await ctx.watch();
    console.log('waiting');
})();