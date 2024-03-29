// 定义一个简单的类 `Person`
import { ConfigGet, ConfigSet } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyHistoryProject } from "@/constant/keys/config";

export class HistoryProject {
  public static currentList: string[] = [];

  // 初始化配置文件中的配置(启动时执行一次即可)
  public static async initList() {
    const currListString = await ConfigGet(ConfigKeyHistoryProject);
    try {
      this.currentList = JSON.parse(currListString);
    } catch (e) {
      this.currentList = [];
    }
  }

  /**
   * 将指定目录添加到当前列表的开头，并确保列表长度不超过10。之后将更新后的列表保存到配置文件中。
   * @param dir 要添加到列表开头的目录路径。
   */
  public static add(dir: string) {
    // 从当前列表中移除指定目录，确保后续添加不会重复
    this.currentList = this.currentList.filter((item) => item !== dir);

    // 将指定目录添加到列表的开头
    this.currentList.unshift(dir);

    // 限制列表长度不超过10，超出部分会被移除
    this.currentList = this.currentList.slice(0, 10);

    // 保存更新后的列表到配置文件
    this.SaveToConfig();
  }

  //保存历史数据到到配置文件
  public static SaveToConfig() {
    ConfigSet(ConfigKeyHistoryProject, JSON.stringify(this.currentList));
  }
}
