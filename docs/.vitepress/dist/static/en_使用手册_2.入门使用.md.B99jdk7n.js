import{_ as e,a as t,b as a,c as o,d as i,e as n,f as s,g as c,h as r,i as l,j as d,k as p,l as h}from"./chunks/29bf6e91-1dfb-4734-be43-10fd29b9dc2a.CMb43z0y.js";import{_ as u,j as m,c as g,o as f,a4 as b}from"./chunks/framework.71dsLfGE.js";const _="_button_1yw4d_2",w={button:_},y=b('<h1 id="introduction-to-software-interface" tabindex="-1">Introduction to Software Interface <a class="header-anchor" href="#introduction-to-software-interface" aria-label="Permalink to &quot;Introduction to Software Interface&quot;">​</a></h1><p><img src="'+e+'" alt="8296bf0dc76e4517a526b42da4378c6a.png"></p><ul><li><ol><li>Navigation bar</li></ol></li><li><ol start="2"><li>Document toolbar</li></ol></li><li><ol start="3"><li>Document directory tree</li></ol></li><li><ol start="4"><li>Document editor area</li></ol></li><li><ol start="5"><li>Basic configuration of the current document</li></ol></li></ul><h1 id="create-a-new-project" tabindex="-1">Create a New Project <a class="header-anchor" href="#create-a-new-project" aria-label="Permalink to &quot;Create a New Project&quot;">​</a></h1><p>Click &quot;Document toolbar - Add Icon - New Project&quot;</p><p><img src="'+t+'" alt="2979d5e6249d4bed8e488b771fc25320.png"></p><p>After clicking, the following interface will pop up:</p><p><img src="'+a+'" alt="5f965c0930e7400a92fe676e7599c0b2.png"></p><blockquote><ol><li>Project name (also the folder name of the project, it is recommended to use English)</li><li>Project description</li><li>Select the directory to store the project (Final project directory: selected directory + project name)</li></ol></blockquote><p>After creating the project, the newly created project will be automatically opened, and you can start document operations as shown in the following figure:</p><p>In the red box is the basic configuration of the currently opened document, which is available for every document.</p><p><img src="'+o+`" alt="22756e46f24047d6a8aad4a464bda7a6.png"></p><h1 id="online-preview" tabindex="-1">Online Preview <a class="header-anchor" href="#online-preview" aria-label="Permalink to &quot;Online Preview&quot;">​</a></h1><ol><li>Go to the project root directory to install dependencies and run the project</li></ol><div class="language- vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki shiki-themes github-light github-dark vp-code"><code><span class="line"><span># Go to the project root directory, not the original directory</span></span>
<span class="line"><span>npm install</span></span>
<span class="line"><span>npm run dosc:dev</span></span></code></pre></div><p><img src="`+i+'" alt="39137c1dddb547a1a9887cfd3ba8be96.png"></p><p>Access the browser: <code>http://localhost:5173</code> to achieve online preview. If everything goes well, you will see the following interface.</p><p><img src="'+n+`" alt="fcc073377585481ba6d05a6b5ee3d01a.png"></p><div class="language-shell vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki shiki-themes github-light github-dark vp-code"><code><span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;"># After editing is completed, use the following command to build</span></span>
<span class="line"><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">npm</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> run</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> dosc:build</span></span></code></pre></div><p><strong>Tip</strong>: You can enter the project&#39;s root directory by clicking on the folder icon in the toolbar.</p><p><img src="`+s+'" alt="b485c946be264eeca0a8ad3f3aea2cf5.png"></p><h2 id="navigation-configuration" tabindex="-1">Navigation Configuration <a class="header-anchor" href="#navigation-configuration" aria-label="Permalink to &quot;Navigation Configuration&quot;">​</a></h2><p>You can edit the top navigation of the website as shown below</p><p><img src="'+c+'" alt="3c63f22b856b4db19b34d3c2e1e40492.png"></p><h1 id="document-management" tabindex="-1">Document Management <a class="header-anchor" href="#document-management" aria-label="Permalink to &quot;Document Management&quot;">​</a></h1><h2 id="creating-a-new-document" tabindex="-1">Creating a New Document <a class="header-anchor" href="#creating-a-new-document" aria-label="Permalink to &quot;Creating a New Document&quot;">​</a></h2><ol><li>Create a new document in the root directory as shown in the following figure:</li></ol><p><img src="'+r+'" alt="7d8c870cbf8b45a1bfb951c380c844cb.png"></p><ol start="2"><li>Create a new document in the specified folder: right-click the folder - New Document</li></ol><p><img src="'+l+'" alt="f291c8dfb59a4c2389d07a6fa514661c.png"></p><h2 id="basic-document-configuration" tabindex="-1">Basic Document Configuration <a class="header-anchor" href="#basic-document-configuration" aria-label="Permalink to &quot;Basic Document Configuration&quot;">​</a></h2><p><img src="'+d+'" alt="c1fa199a8271473e899f635df0f7d494.png"></p><p>As shown in the above figure, you can configure individual settings for the document.</p><p>The following fields are specifically explained:</p><p><strong>Outline Display Level</strong>: When set to 0, the outline is not displayed; when set to 3, the outline can be displayed up to the third level.</p><h2 id="document-as-homepage" tabindex="-1">Document as Homepage <a class="header-anchor" href="#document-as-homepage" aria-label="Permalink to &quot;Document as Homepage&quot;">​</a></h2><p>If you want the previously edited document to be the homepage, you can click the &quot;Homepage&quot; tab on the right side, and select &quot;Homepage&quot; as the page type.</p><p>Configure the basic information and features of this homepage.</p><p><img src="'+p+'" alt="95c8f4e20e604dc79fabcc86352a858e.png"></p><h2 id="document-code" tabindex="-1">Document Code <a class="header-anchor" href="#document-code" aria-label="Permalink to &quot;Document Code&quot;">​</a></h2><blockquote><p>In VitePress, each Markdown file is compiled into HTML and treated as a Vue single-file component. This means that you can use any Vue features in Markdown, including dynamic templates, Vue components, or adding logic to the Vue component of the page by adding <code>&lt;script&gt;</code> tags.</p></blockquote><p>Code refers to the code in the script and style tags of the document. The program will automatically recognize the code in these two tags and fill it into the code area (the entire document should not have multiple script or style tags, as unexpected events may occur during parsing).</p><p><img src="'+h+'" alt="29bf6e911dfb4734be4310fd29b9dc2a.png"></p>',43),k=[y],N=JSON.parse('{"title":"这是文档的标题","description":"这是文档的seo描述å","frontmatter":{"title":"这是文档的标题","description":"这是文档的seo描述å","navbar":true,"sideBar":true,"footer":false,"outline":[1,3],"editLink":true,"lastUpdated":true,"aside":"right","layout":"doc","custom":{"A1":"Hhh","":""},"hero":{"image":{"src":"","alt":"","width":"","height":""},"name":"VitePressSimple","text":"quick to config vitePress","description":"","tagline":"","actions":[],"features":[],"head":[]},"head":[["meta",{"name":"Author","content":"Xiaod"}],["meta",{"name":"","content":""}]]},"headers":[],"relativePath":"en/使用手册/2.入门使用.md","filePath":"en/使用手册/2.入门使用.md"}'),v={name:"en/使用手册/2.入门使用.md"},q=Object.assign(v,{setup(P){return m(0),(j,D)=>(f(),g("div",null,k))}}),C={$style:w},I=u(q,[["__cssModules",C]]);export{N as __pageData,I as default};
