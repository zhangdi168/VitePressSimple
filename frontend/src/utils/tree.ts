// 工具类
import { DataNode } from "ant-design-vue/es/tree";

export interface VpNav {
  link: string;
  text: string;
  items?: VpNav[];
}

export interface VpSidebar {
  link: string;
  text: string;
  items?: VpSidebar[];
}

// 将VpNav类型数组转换为DataNode类型数组
export function convertVpNavArrayToDataNode(vpNavArray: VpNav[]): DataNode[] {
  return vpNavArray.map((vpNav) => ({
    key: vpNav.link ?? "",
    title: vpNav.text ?? "",
    children: vpNav.items ? vpNav.items.map(convertVpNavToDataNode) : undefined,
  }));
}

// 将VpNav类型转换为DataNode类型
function convertVpNavToDataNode(vpNav: VpNav): DataNode {
  return {
    key: vpNav.link ?? "",
    title: vpNav.text ?? "",
    children: vpNav.items ? vpNav.items.map(convertVpNavToDataNode) : undefined,
  };
}

// 将DataNode类型数组转换为VpNav类型数组
export function convertDataNodeArrayToVpNav(dataNodeArray: DataNode[]): any[] {
  return dataNodeArray.map((dataNode) => {
    if (!dataNode.children) {
      // 如果没有children，则只返回text和link部分
      return {
        text: dataNode.title ?? "",
        link: dataNode.key,
      };
    } else {
      // 如果有children，则只返回text和link部分
      return {
        text: dataNode.title ?? "",
        items: dataNode.children.map(convertDataNodeToVpNav),
      };
    }
  });
}

// 将DataNode类型转换为VpNav类型
function convertDataNodeToVpNav(dataNode: DataNode): any {
  if (!dataNode.children) {
    // 如果没有children，则只返回text和link部分
    return {
      text: dataNode.title ?? "",
      link: dataNode.key,
    };
  } else {
    // 如果有children，则只返回text和link部分
    return {
      text: dataNode.title ?? "",
      items: dataNode.children.map(convertDataNodeToVpNav),
    };
  }
}

// 删除节点
// TreeUtil.deleteNodeByKey(data, 'someKey');
//
//  修改节点
// TreeUtil.updateNodeByKey(data, 'someKey', { title: 'newTitle' });
//  新增子节点
// const newNode: DataNode = { key: 'newKey', title: 'newTitle' };
export class TreeUtil {
  /**
   * 根据key查找节点
   * @param {DataNode[]} data - 树形数据
   * @param {string | number} key - 要查找的key
   * @returns {DataNode | undefined} - 找到的节点，未找到则返回undefined
   */
  static findNodeByKey(data: DataNode[], key: string | number): DataNode {
    for (const node of data) {
      if (node.key === key) return node;
      if (node.children) {
        const foundNode = this.findNodeByKey(node.children, key);
        if (foundNode) return foundNode;
      }
    }
    return {
      key: "undefined",
      title: "undefined",
      children: [],
    };
  }

  /**
   * 删除指定key的节点
   * @param {DataNode[]} data - 树形数据
   * @param {string | number} key - 要删除的节点key
   * @returns {boolean} - 是否成功删除
   */
  static deleteNodeByKey(data: DataNode[], key: string | number): boolean {
    for (let i = 0; i < data.length; i++) {
      if (data[i].key === key) {
        data.splice(i, 1);
        return true;
      } else if (data[i].children) {
        // @ts-ignore
        if (this.deleteNodeByKey(data[i].children, key)) {
          return true;
        }
      }
    }
    return false;
  }

  /**
   * 修改指定key的节点
   * @param {DataNode[]} data - 树形数据
   * @param {string | number} key - 要修改的节点key
   * @param {Partial<DataNode>} updateInfo - 需要更新的信息
   * @returns {boolean} - 是否成功修改
   */
  static updateNodeByKey(
    data: DataNode[],
    key: string | number,
    updateInfo: Partial<DataNode>,
  ): boolean {
    for (const node of data) {
      if (node.key === key) {
        Object.assign(node, updateInfo);
        return true;
      } else if (node.children) {
        if (this.updateNodeByKey(node.children, key, updateInfo)) return true;
      }
    }
    return false;
  }

  /**
   * 在指定key的节点下新增子节点
   * @param {DataNode[]} data - 树形数据
   * @param {string | number} parentKey - 新增子节点的父节点key
   * @param {DataNode} newNode - 新增的子节点
   * @returns {boolean} - 是否成功新增
   */
  static addNodeToParentByKey(
    data: DataNode[],
    parentKey: string | number,
    newNode: DataNode,
  ): boolean {
    for (const node of data) {
      if (node.key === parentKey) {
        if (!node.children) node.children = [];
        node.children.push(newNode);
        return true;
      } else if (node.children) {
        if (this.addNodeToParentByKey(node.children, parentKey, newNode))
          return true;
      }
    }
    return false;
  }
}
