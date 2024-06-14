export namespace dto {
	
	export class TreeNode {
	    title: string;
	    key: string;
	    style?: string;
	    class?: string;
	    children?: TreeNode[];
	
	    static createFrom(source: any = {}) {
	        return new TreeNode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.key = source["key"];
	        this.style = source["style"];
	        this.class = source["class"];
	        this.children = this.convertValues(source["children"], TreeNode);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace services {
	
	export class NotifyShellData {
	    shell_index?: number;
	    content?: string;
	    is_error?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NotifyShellData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.shell_index = source["shell_index"];
	        this.content = source["content"];
	        this.is_error = source["is_error"];
	    }
	}

}

