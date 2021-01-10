const path = require('path');
const webpack = require('webpack');
const CopyPlugin = require('copy-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
require("@babel/polyfill");

module.exports = {
  entry: ['@babel/polyfill', path.resolve(__dirname, 'app', 'index.tsx')],
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
    publicPath: '/'
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new CopyPlugin({
      patterns: [
        {
          from: 'app/assets',
          to: 'assets'
        }
      ]
    }),
    new HtmlWebpackPlugin({
      template: 'index.html'
    })
  ],
  module: {
    rules: [
      {
        test:/\.(s*)css$/,
        use: ['style-loader', 'css-loader', 'sass-loader'],
        include: path.resolve(__dirname, 'app'),
      },
      {
        test: /\.(jsx|js|ts|tsx)$/,
        include: path.resolve(__dirname, 'app'),
        exclude: /node_modules/,
        use: [{
          loader: 'babel-loader',
          options: {
            presets: [
              [
                '@babel/preset-env', { targets: "defaults" }
              ],
              '@babel/preset-react',
              '@babel/preset-typescript'
            ]
          }
        }]
      }
    ]
  },
  devServer: {
    contentBase: path.resolve(__dirname, 'dist'),
    open: false,
    clientLogLevel: 'silent',
    port: 4200,
    historyApiFallback: true,
    index: 'index.html',
    proxy: {
      '/api': {
        target: 'http://localhost:5775',
        secure: false
      }
    }
  },
  resolve: {
    extensions: ['.ts', '.js', '.jsx', '.tsx']
  }
}
