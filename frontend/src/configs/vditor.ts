export const getDefaultVtitorOptions = (): IOptions => {
  return {
    // cdn: "https://fastly.jsdelivr.net/npm/",
    //  cdn: "http://localhost:9874/cdn",
    height: "90vh",
    toolbarConfig: {
      pin: true,
    },
    counter: {
      enable: true,
      type: "text",
    },
    value: ``,
    toolbar: [
      // 表情符号插入按钮
      "emoji",
      // 标题级别选择按钮（例如 H1, H2, H3 等）
      "headings",
      // 加粗文本按钮
      "bold",
      // 斜体文本按钮
      "italic",
      // 删除线文本按钮
      "strike",
      // 分隔符，用于分隔不同功能区域
      "|",
      // 插入水平分割线按钮
      "line",
      // 引用块按钮
      "quote",
      // 无序列表按钮
      "list",
      // 有序列表按钮
      "ordered-list",
      // 复选框列表项
      "check",
      // 减少缩进按钮
      "outdent",
      // 增加缩进按钮
      "indent",
      // 插入代码块按钮
      "code",
      // 插入行内代码按钮
      "inline-code",
      // 在当前光标位置之后插入内容
      // "insert-after",
      // // 在当前光标位置之前插入内容
      // "insert-before",
      // 撤销操作按钮
      "undo",
      // 重做操作按钮
      "redo",
      // 文件上传按钮
      // 'upload',
      // 插入链接按钮
      "link",
      // 插入表格按钮
      "table",
      // 录音按钮（如果支持的话）
      // 'record',
      // 编辑模式切换按钮（例如：编辑/预览模式）
      "edit-mode",
      // 同时显示编辑和预览区域的模式
      "both",
      // 预览模式按钮
      // "preview",
      // 全屏模式按钮
      // 'fullscreen',
      // 文档大纲按钮，展示文档结构
      "outline",
      // 代码主题切换按钮
      "code-theme",
      // 内容主题切换按钮
      "content-theme",
      // 导出内容按钮
      // "export",
      // 开发者工具（如需调试或查看内部信息）
      // 'devtools',
      // 显示关于编辑器的信息
      // 'info',
      // 显示帮助信息或快捷键指南
      // 'help',
      // 插入换行符按钮
      // 'br'
    ],
    cache: {
      enable: false,
    },
    outline: {
      enable: false,
      position: "right",
    },
    upload: {
      url: "",
      max: 1,
    },
  };
};
