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

// corsエラーになるためOpenSearch直アクセスではなく、API経由で実行できるようにする
async function searchComment() {
  const product_id = document.getElementById('search_product_id').value;
  const comment_id = document.getElementById('search_comment_id').value;

  const query = {
      query: {
          bool: {
              must: [
                  { match: { product_id: Number(product_id) } },
                  { term: { _id: comment_id } }
              ]
          }
      }
  };

  const response = await fetch("http://localhost:9200/product_comments/_search?pretty", {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(query)
  });

  const result = await response.json();
  document.getElementById('search_result').textContent = JSON.stringify(result, null, 2);
}
