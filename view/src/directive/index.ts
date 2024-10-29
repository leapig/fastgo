import type { App } from 'vue'
import pagePermi from './permission/pagePermi'
import readPermi from './permission/readPermi'
import writePermi from './permission/writePermi'

export default function directive(app: App<Element>){
  app.directive('pagePermi', pagePermi)
  app.directive('readPermi', readPermi)
  app.directive('writePermi', writePermi)
}
