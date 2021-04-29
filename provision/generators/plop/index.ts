import { NodePlopAPI } from 'node-plop'
import { testGenerator } from './generators'
import { sanitize } from './helpers'

export default function plop(plop: NodePlopAPI) {
  plop.setGenerator('test', testGenerator)
  plop.setHelper('sanitize', sanitize)
}
