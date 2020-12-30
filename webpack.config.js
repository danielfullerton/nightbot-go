const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const webpack = require('webpack');

module.exports = {
  entry: path.resolve(__dirname, 'app', 'index.tsx'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: "index.html",
      template: "app/index.html"
    }),
    new webpack.HotModuleReplacementPlugin()
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
    open: true,
    clientLogLevel: 'silent',
    port: 4200
  },
  resolve: {
    extensions: ['.ts', '.js', '.jsx', '.tsx']
  }
}
