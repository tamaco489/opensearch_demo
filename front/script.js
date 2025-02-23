async function postComment() {
  const product_id = document.getElementById('product_id').value;
  const title = document.getElementById('title').value;
  const content = document.getElementById('content').value;
  const rate = document.getElementById('rate').value;

  const response = await fetch(`http://localhost:8080/shop/v1/products/${product_id}/comments`, {
      method: 'POST',
      headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({ title, content, rate: Number(rate) })
  });

  const result = await response.json();
  document.getElementById('post_result').textContent = `投稿成功！コメントID: ${result.id}`;
}

async function getProductCommentByID() {
  const product_id = document.getElementById('get_product_comment_id').value;
  const comment_id = document.getElementById('get_comment_id').value;

  console.log(`📝 取得した商品ID: ${product_id}`);
  console.log(`📝 取得したコメントID: ${comment_id}`);

  if (!product_id || !comment_id) {
    console.error("⚠️ 商品IDとコメントIDが空です！");
    document.getElementById('get_comment_result').textContent = "商品IDとコメントIDを入力してください。";
    return;
  }

  const url = `http://localhost:8080/shop/v1/products/${product_id}/comments/${comment_id}`;
  console.log(`📡 送信するリクエストURL: ${url}`);

  try {
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Accept': 'application/json'
      }
    });

    console.log(`📡 HTTPステータスコード: ${response.status}`);

    if (!response.ok) {
      throw new Error(`HTTPエラー: ${response.status}`);
    }

    const result = await response.json();
    console.log("✅ API レスポンス:", result);

    document.getElementById('get_comment_result').textContent = JSON.stringify(result, null, 2);
  } catch (error) {
    console.error("❌ エラー:", error);
    document.getElementById('get_comment_result').textContent = `エラー: ${error.message}`;
  }
}
