<script setup lang="ts">
import { ConfigGet, OpenURL } from "../../../wailsjs/go/system/SystemService";
import { useIndexStore } from "../../store";
import { onMounted } from "vue";
import { ConfigKeyLang } from "@/constant/keys/config";

let lang = "";
onMounted(async () => [(lang = await ConfigGet(ConfigKeyLang))]);
</script>

<template>
  <div v-if="lang == 'en'">
    <div class="mx-20 mt-4 flex justify-center border p-3">
      <span class="text-3xl font-bold text-center"
        >About VPSimple {{ useIndexStore().version }}</span
      >
      <div class="px-4">
        <p>
          This software is built with Vue3, TypeScript, and Wails, aiming to
          provide VitePress users with a convenient and visual configuration
          experience.
        </p>

        <ul>
          <li>
            Named "VitePressSimple" (short for VPSimple), the source code of
            this software is made fully open-source on
            <span
              class="cursor-pointer"
              @click="OpenURL('https://github.com/zhangdi168/VitePressSimple')"
              ><strong>GitHub</strong></span
            >
            under the <strong>Apache 2.0</strong> license.
          </li>
          <li>
            Any individual or organization can freely access, download, use,
            study, modify, and redistribute this software, as long as they
            adhere to the corresponding open-source license terms.
          </li>
        </ul>

        <p>
          "<strong>VitePressSimple</strong>" focuses on enhancing the efficiency
          and user experience of configuring VitePress projects with the
          following core features:
        </p>

        <ul>
          <li>
            <strong>Developed with Vue3 and TypeScript:</strong>
            Leveraging the powerful component-based system of Vue3 and the
            static type checking of TypeScript, ensuring high performance,
            stability, and maintainability.
          </li>
          <li>
            <strong>Integrated with Wails framework:</strong>
            Built as a desktop client application, providing a seamless
            connection to the user's local environment with a native, intuitive
            operation experience.
          </li>
          <li>
            <strong>Quick configuration:</strong>
            Offers a rich set of VitePress configuration options through a clean
            interface, simplifying project initialization and advanced
            configurations, reducing manual setup efforts.
          </li>
          <li>
            <strong>Visual configuration interface:</strong>
            Provides an easy-to-understand graphical interface that demonstrates
            configuration effects in real-time, making complex VitePress
            configurations straightforward and minimizing the learning curve and
            error probability.
          </li>
        </ul>

        <p>
          We warmly welcome you to try "<strong>VitePressSimple</strong>" and
          look forward to your valuable feedback and suggestions. You can
          interact with us through the following channels:
        </p>

        <ul>
          <li>
            <strong>Report issues or suggest features:</strong>
            Submit issues or feature requests on the open-source platform's
            Issue section, and we will respond promptly and consider
            incorporating useful suggestions.
          </li>
          <li>
            <strong>Contribute code and documentation:</strong>
            If you'd like to contribute to the software's development and
            improvement, feel free to submit a Pull Request.
          </li>
          <li>
            <strong>User discussion and interaction:</strong>
            Join our
            <a href="[Community Link (e.g., Discord, forum)]"
              ><strong>GitHub community</strong></a
            >
            to share experiences, discuss technical issues with other users and
            developers.
          </li>
        </ul>

        <p>
          While we strive to ensure the quality and suitability of the software,
          due to its open-source nature and user-configurable flexibility, we
          cannot be held responsible for any direct, indirect losses, or issues
          arising in special circumstances resulting from the use of this
          software. Users should abide by relevant laws and regulations when
          using the software and must not use it for any illegal, infringing, or
          detrimental purposes.
        </p>
      </div>
    </div>
  </div>
  <div v-else>
    <div class="mx-20 mt-4 flex justify-center border p-3">
      <span class="text-3xl font-bold text-center"
        >关于VPSimple {{ useIndexStore().version }}</span
      >
      <div class="px-4">
        <p>
          本软件以Vue3、TypeScript与Wails技术栈为核心，旨在为VitePress用户提供便捷与可视化配置体验。
        </p>

        <ul>
          <li>
            本软件命名为“VitePressSimple”(简称VPSimple)，遵循开放源代码的原则，我们已将其源代码在
            <span
              class="cursor-pointer"
              @click="OpenURL('https://gitee.com/zhangdi168/VitePressSimple')"
              ><strong>github</strong></span
            >上完全公开，遵循<strong>Apache 2.0</strong>。
          </li>
          <li>
            任何个人或组织皆可自由访问、下载、使用、研究、修改及再分发本软件，只需遵守相应开源许可协议的规定。
          </li>
        </ul>

        <p>
          “<strong>VitePressSimple</strong>”专注于提升VitePress项目的配置效率与用户体验，其核心特性如下：
        </p>

        <ul>
          <li>
            <strong>基于Vue3与TypeScript开发：</strong
            >利用Vue3的强大组件化体系与TypeScript的静态类型检查能力，确保软件具备出色的性能、稳定性与可维护性。
          </li>
          <li>
            <strong>集成Wails框架：</strong
            >构建为桌面客户端应用，无缝衔接用户本地环境，提供直观、流畅的原生操作体验。
          </li>
          <li>
            <strong>快速配置功能：</strong
            >预设丰富的VitePress配置选项，用户只需通过简洁明了的界面即可完成项目初始化与各项高级配置，极大地简化了繁琐的手动配置过程。
          </li>
          <li>
            <strong>可视化配置界面：</strong
            >采用直观易用的图形化界面，实时展现配置效果，使复杂的VitePress配置变得一目了然，显著降低用户的学习成本与出错概率。
          </li>
        </ul>

        <p>
          我们诚挚欢迎广大用户试用“<strong>VitePressSimple</strong>”，并期待您的宝贵反馈与建议。您可以通过以下途径与我们互动：
        </p>

        <ul>
          <li>
            <strong>问题报告与功能建议：</strong
            >在开源平台的Issue板块提交问题或提出新功能建议，我们将及时响应并积极采纳有益意见。
          </li>
          <li>
            <strong>贡献代码与文档：</strong
            >如果您愿意参与软件的进一步开发与完善，欢迎您提交Pull
            Request，共同提升软件品质。
          </li>
          <li>
            <strong>用户交流与讨论：</strong>加入我们的<a
              href="[社区链接（如Discord、论坛等）]"
              ><strong>github社区</strong></a
            >，与其他用户、开发者共享使用经验，探讨技术问题。
          </li>
        </ul>

        <p>
          尽管我们竭力保证软件的质量与适用性，但鉴于软件的开源性质与用户自行配置的灵活性，我们无法对因使用本软件而导致的任何直接、间接损失或特殊情况下引发的问题承担责任。用户在使用过程中应遵守相关法律法规，不得将本软件用于任何违法、侵权或损害公共利益的目的。
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
