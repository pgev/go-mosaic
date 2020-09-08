/* eslint-disable @typescript-eslint/no-var-requires */

const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  target: 'web',

  entry: path.join(__dirname, 'src', 'index.tsx'),

  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
  },

  module: {
    rules: [
      {
        test: /\.js$/,
        enforce: 'pre',
        exclude: /node_modules/,
        use: ['source-map-loader'],
      },
      {
        test: /\.(js|ts)x?$/,
        exclude: /node_modules/,
        use: [{
          loader: 'babel-loader',
          query: {
          },
        }],
      },
      {
        exclude: /node_modules/,
        test: /\.graphql$/,
        use: [{ loader: 'graphql-import-loader' }],
      },
    ],
  },

  output: {
    filename: 'bundle.js',
    path: path.join(__dirname, 'www'),
  },

  plugins: [
    new HtmlWebpackPlugin({
      template: path.join(__dirname, 'index.html'),
    }),
  ],

  devtool: 'source-map',

  optimization: {
    minimize: true,
  },

  devServer: {
    contentBase: path.join(__dirname, 'www'),
    compress: true,
    port: 9000,
  },
};
