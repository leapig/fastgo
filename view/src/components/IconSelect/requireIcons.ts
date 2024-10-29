import * as ElIcon from '@element-plus/icons-vue'
import FontIcon from '@/assets/fonts/iconfont/iconfont.json'

export const defaultType = 'elIcons'

export const types = [
  { title: 'el-icon', key: 'elIcons' },
  { title: 'iconfont', key: 'fontIcons' },
]

const thisIcons = {
  elIcons: [] as string[],
  fontIcons: [] as string[]
}

// 加载 element-plus-icons
thisIcons.elIcons = Object.keys(ElIcon)
// 加载 iconfont-icons
if (FontIcon && FontIcon.glyphs) {
  FontIcon.glyphs.forEach((item: any) => {
    thisIcons.fontIcons.push(FontIcon.css_prefix_text + item.font_class)
  })
}

export const icons = thisIcons
