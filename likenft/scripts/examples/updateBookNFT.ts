import { ethers } from "hardhat";

async function updateBookNFT() {
  const classId = "0x1D146390C1D4E03C74b87D896b254a5468EDF804";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("BookNFT", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  const tx = await likeNFTClass.update({
    name: "《所謂「我不投資」，就是 all in 在法定貨幣》2",
    symbol: "BOOK",
    metadata: JSON.stringify({
      name: "《所謂「我不投資」，就是 all in 在法定貨幣》2",
      symbol: "BOOK",
      description:
        "購買本書會透過郵件收到 epub 和 pdf 檔，同時得到換領本書 NFT 的資格。書的 epub 和 pdf 歡迎轉發給朋友，讓更多人獲得相關知識。有別於粵語中「翻版」的原意，「翻版」不是「盜版」，並不違法。\n然而，翻版雖然不是盜版，卻也不是正版。唯有 NFT 的持有者，手上的 epub 和 pdf 才是正版。購買正版是一份美德，是能力所及的讀者該付出的一點承擔，代表著對作者的支持，對知識、報道和創作的尊重。\n翻版是傳播，正版是應援；傳播的同時，不要忘了鼓勵對方購買正版，讓正念 pay it forward。\n鼓勵正版，允許翻版，消滅盜版，是《所謂「我不投資」，就是 all in 在法定貨幣》分散式出版的核心理念。",
      image:
        "ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey",
      banner_image: "",
      featured_image: "",
      external_link: "https://bit.ly/moneyverse-pdf",
      collaborators: [],
    }),
    max_supply: 20,
  });
  console.log(await tx.wait());
}

updateBookNFT().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
