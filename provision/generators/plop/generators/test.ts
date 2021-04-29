import { Actions, PlopGeneratorConfig } from 'node-plop'
import slugify from 'slugify'
import * as path from 'path'
import { ImagePrompNames, Answers } from './entities'
import { baseRootPath, baseTemplatesPath } from '../utils'
const testPath = path.join(baseRootPath, 'test')

export const testGenerator: PlopGeneratorConfig = {
  description: 'add an path to test',
  prompts: [
    {
      type: 'input',
      name: ImagePrompNames.name,
      message: 'What should it be image?',
      default: 'youtubedl'
    },
    {
      type: 'confirm',
      name: ImagePrompNames.wantTest,
      default: true,
      message: 'Do you want implement test?'
    }
  ],
  actions: (data) => {
    const answers = data as Answers

    const actions: Actions = []

    actions.push({
      type: 'append',
      templateFile: `${baseTemplatesPath}/image/Dockerfile.add.hbs`,
      path: `${baseRootPath}/${slugify(answers.imageName, '_')}/Dockerfile`,
      abortOnFail: true
    })

    if (answers.wantTest) {
      actions.push({
        type: 'append',
        templateFile: `${baseTemplatesPath}/test/test.add.hbs`,
        path: `${testPath}/docker_${slugify(answers.imageName, '_')}_test.go`,
        abortOnFail: true
      })
    }

    return actions
  }
}
