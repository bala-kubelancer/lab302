package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	appName    = getEnv("APP_NAME", "Lab302")
	appVersion = getEnv("APP_VERSION", "v2") // ✅ v2
	appEnv     = getEnv("APP_ENV", "dev")
	listenAddr = ":8080"
	startTime  = time.Now()
	requests   = 0
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func handler(w http.ResponseWriter, r *http.Request) {
	requests++
	uptime := time.Since(startTime).Round(time.Second)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>%s – AKS GitOps Demo</title>

<style>
:root {
  --accent:#22c55e; /* ✅ v2 accent */
  --bg1:#020617;
  --bg2:#020617;
}

* { margin:0; padding:0; box-sizing:border-box; }

body {
  min-height:100vh;
  display:flex;
  flex-direction:column;
  align-items:center;
  justify-content:center;
  background:linear-gradient(180deg, var(--bg1), var(--bg2));
  font-family:Inter,system-ui;
  color:#e5e7eb;
}

header, footer {
  width:100%%;
  text-align:center;
  padding:18px 0;
  letter-spacing:.35em;
  font-size:12px;
  text-transform:uppercase;
  opacity:.6;
}

/* ===== v2 Badge ===== */
.badge {
  display:inline-block;
  margin-top:10px;
  padding:6px 16px;
  border-radius:999px;
  background:linear-gradient(135deg, #22c55e, #86efac);
  color:#020617;
  font-weight:900;
  font-size:13px;
}

.glow {
  position:absolute;
  width:900px;
  height:900px;
  background:radial-gradient(circle, var(--accent), transparent 70%%);
  filter:blur(150px);
  opacity:.35;
  z-index:-1;
}

.card {
  width:820px;
  padding:56px;
  border-radius:32px;
  background:rgba(15,23,42,.95);
  border:1px solid rgba(255,255,255,.08);
  box-shadow:0 80px 160px rgba(0,0,0,.8);
}

.title {
  display:flex;
  justify-content:space-between;
  align-items:center;
  margin-bottom:30px;
}

.name { font-size:40px; font-weight:800; }

.version {
  font-size:22px;
  padding:14px 34px;
  border-radius:999px;
  background:linear-gradient(135deg, var(--accent), white);
  color:black;
  font-weight:900;
}

.status {
  display:flex;
  align-items:center;
  font-size:18px;
  margin-bottom:30px;
}

.dot {
  width:16px;
  height:16px;
  border-radius:999px;
  background:#22c55e;
  margin-right:12px;
  box-shadow:0 0 18px #22c55e;
}

.grid {
  display:grid;
  grid-template-columns:1fr 1fr;
  gap:24px;
}

.box {
  padding:22px;
  border-radius:18px;
  background:rgba(255,255,255,.05);
  border:1px solid rgba(255,255,255,.08);
}

.label {
  font-size:12px;
  letter-spacing:.2em;
  text-transform:uppercase;
  opacity:.6;
  margin-bottom:10px;
}

.value {
  font-size:20px;
  font-family:JetBrains Mono, monospace;
  color:#5eead4;
}
</style>
</head>

<body>

<header>
  <h1>AKS GITOPS DEMO</h1>
  <div class="badge">VERSION v2</div>
  <div style="margin-top:6px; font-size:12px; opacity:.6;">
    ARGO CD · DECLARATIVE DEPLOYMENTS
  </div>
</header>

<div class="glow"></div>

<div class="card">

  <div class="title">
    <div class="name">%s</div>
    <div class="version">%s</div>
  </div>

  <div class="status">
    <div class="dot"></div>
    Released and operational
  </div>

  <div class="grid">
    <div class="box">
      <div class="label">Environment</div>
      <div class="value">%s</div>
    </div>
    <div class="box">
      <div class="label">Requests / Uptime</div>
      <div class="value">%d · %s</div>
    </div>
  </div>

</div>

<footer>
  <div>Powered by</div>
  <h1>KUBELANCER LABS</h1>
</footer>

</body>
</html>`,
		appName,
		appName,
		appVersion,
		appEnv,
		requests,
		uptime,
	)

	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("🚀 %s %s (%s)\n", appName, appVersion, appEnv)
	http.ListenAndServe(listenAddr, nil)
}
