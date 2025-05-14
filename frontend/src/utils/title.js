// 更新页面标题
export function updateTitle(title) {
  document.title = title
}

// 更新页面标题，带有子页面标题
export function updatePageTitle(systemName, pageTitle) {
  document.title = pageTitle ? `${pageTitle} - ${systemName}` : systemName
} 