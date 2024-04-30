use anyhow::Result;

#[tokio::test]
async fn quick_dev() -> Result<()> {
    let hc = httpc_test::new_client("http://localhost:3003")?;
    hc.do_get("/health").await?.print().await?;
    hc.do_get("/info").await?.print().await?;
    Ok(())
}