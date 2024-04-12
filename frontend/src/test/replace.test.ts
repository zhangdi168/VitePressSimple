import { test } from "vitest";
import { replaceLocalStaticToImageUrl } from "@/utils/repalceStatic";

test("repalce1", () => {
  const content =
    "![d8e8a54a8b8b491dbd5eae2495898a51.png](/vpstatic/images/20240411/d8e8a54a-8b8b-491d-bd5e-ae2495898a51.png)\n";
  const newVal = replaceLocalStaticToImageUrl(content);
  console.log(newVal, "newVal -- console.log");
});
