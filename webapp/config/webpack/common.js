// shared config (dev and prod)
const webpack = require('webpack');
const { resolve } = require("path");
const Dotenv = require("dotenv-webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const CopyPlugin = require("copy-webpack-plugin");


module.exports = {
  target: 'web',
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"],
    fallback: {
      buffer: require.resolve('buffer'),
    },
  },
  context: resolve(__dirname, "../../src"),
  module: {
    rules: [
      {
        test: [/\.jsx?$/, /\.tsx?$/],
        use: ["babel-loader"],
        exclude: /node_modules/,
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"],
      },
      {
        test: /\.(scss|sass)$/,
        use: ["style-loader", "css-loader", "sass-loader"],
      },
      {
        test: /\.(jpe?g|png|gif)$/i,
        use: [
          "file-loader?hash=sha512&digest=hex&name=img/[contenthash].[ext]",
          "image-webpack-loader?bypassOnDebug&optipng.optimizationLevel=7&gifsicle.interlaced=false",
        ],
      },
      {
        test: /\.svg$/i,
        issuer: /\.[jt]sx?$/,
        use: ['@svgr/webpack'],
      },
    ],
  },
  plugins: [
    new webpack.ProvidePlugin({
      Buffer: ['buffer', 'Buffer'],
    }),
    new Dotenv({
      systemvars: true,
    }),
    new HtmlWebpackPlugin({ template: resolve('public', 'index.html'), favicon: resolve('public', 'favicon.ico') }),
    new CopyPlugin({
      patterns: [
        { from: resolve('public', 'static'), to: resolve('dist', 'static') },
        { from: resolve('public', 'site.webmanifest'), to: resolve('dist', 'site.webmanifest') },
      ],
    }),
  ],
  performance: {
    hints: false,
  },
};
