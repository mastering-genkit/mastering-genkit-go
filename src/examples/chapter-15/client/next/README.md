# Cooking Battle Client - Next.js 15

Genkit Go サーバーと通信するクライアントアプリケーション。Onion Architecture を採用し、Next.js 15 で実装されています。

## アーキテクチャ

### Onion Architecture

このプロジェクトは Onion Architecture に基づいて構成されています：

```
src/
├── domain/               # ドメイン層（最内層）
│   ├── models/          # ドメインモデル
│   │   ├── chat/        # チャット関連モデル
│   │   ├── action/      # アクション関連モデル
│   │   ├── recipe/      # レシピ関連モデル
│   │   └── error/       # エラー関連モデル
│   └── repositories.ts  # リポジトリインターフェース
│
├── usecases/            # ユースケース層
│   ├── chat.ts         # チャットのユースケース
│   └── action.ts       # アクションのユースケース
│
├── infrastructure/      # インフラストラクチャ層（最外層）
│   ├── http/           # HTTP通信の実装
│   │   ├── dto/        # Data Transfer Objects
│   │   ├── mappers/    # DTO ↔ ドメインモデルのマッピング
│   │   ├── client/     # HTTPクライアントユーティリティ
│   │   ├── config/     # HTTP設定
│   │   └── repository/ # リポジトリの実装
│   └── auth/           # 認証関連
│       └── firebase.ts # Firebase Anonymous Auth
│
├── app/                 # プレゼンテーション層
│   ├── composition.ts   # DI（依存性注入）設定
│   ├── hooks/          # React カスタムフック
│   └── chat/           # チャットページ
│
└── components/          # UIコンポーネント
    ├── MessageList.tsx  # メッセージ一覧
    ├── Composer.tsx     # メッセージ入力
    └── ActionPanel.tsx  # アクションボタン
```

### 層の責務

1. **ドメイン層**
   - ビジネスロジックとドメインモデルを定義
   - 外部依存を持たない純粋なTypeScript
   - リポジトリインターフェースの定義

2. **ユースケース層**
   - アプリケーションのビジネスロジック
   - ドメイン層のみに依存
   - リポジトリインターフェースを通じてデータアクセス

3. **インフラストラクチャ層**
   - 外部システムとの通信実装
   - HTTPクライアント、認証、DTOマッピング
   - リポジトリインターフェースの実装

4. **プレゼンテーション層**
   - UI コンポーネントとページ
   - ユースケースを通じてビジネスロジックを実行
   - React/Next.js 固有の実装

## セットアップ

### 前提条件

- Node.js 18以上
- npm または yarn
- Genkit Go サーバーが起動していること（http://127.0.0.1:9090）

### インストール

```bash
# 依存関係のインストール
npm install

# 環境変数の設定
cp .env.local.example .env.local
# .env.local を編集して必要な設定を追加
```

### 起動方法

```bash
# 開発サーバーの起動
npm run dev

# ブラウザで http://localhost:3000 を開く
```

## 開発ガイド

### 新しい機能の追加

1. **ドメインモデルの追加**
   ```typescript
   // src/domain/models/[feature]/[model].ts
   export interface NewFeature {
     id: string;
     // ... プロパティを定義
   }
   ```

2. **リポジトリインターフェースの定義**
   ```typescript
   // src/domain/repositories.ts
   export interface NewFeatureRepository {
     getFeature(id: string): Promise<NewFeature>;
   }
   ```

3. **ユースケースの実装**
   ```typescript
   // src/usecases/new-feature.ts
   export class GetNewFeatureUseCase {
     constructor(private repository: NewFeatureRepository) {}
     
     async execute(id: string): Promise<NewFeature> {
       return this.repository.getFeature(id);
     }
   }
   ```

4. **インフラストラクチャの実装**
   ```typescript
   // src/infrastructure/http/repository/new-feature-repo.ts
   export class HttpNewFeatureRepository implements NewFeatureRepository {
     async getFeature(id: string): Promise<NewFeature> {
       // HTTP通信の実装
     }
   }
   ```

5. **DIの設定**
   ```typescript
   // app/composition.ts
   const newFeatureRepo = new HttpNewFeatureRepository();
   export const getNewFeature = new GetNewFeatureUseCase(newFeatureRepo);
   ```

### Genkit Flow との通信

Genkit Flow は特別なリクエスト/レスポンス形式を使用します：

1. **リクエスト**：`{ data: <payload> }` でラップ
2. **レスポンス**：
   - 通常：`{ result: <data> }`
   - ストリーミング：`{ message: <chunk> }` または `{ result: <final> }`

### 環境変数

`.env.local` で以下の変数を設定：

```env
# API設定
NEXT_PUBLIC_API_BASE_URL=http://127.0.0.1:9090

# Firebase設定（開発環境ではプレースホルダーでOK）
NEXT_PUBLIC_FIREBASE_API_KEY=placeholder
NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=placeholder
NEXT_PUBLIC_FIREBASE_PROJECT_ID=placeholder
# ... その他のFirebase設定
```

### デバッグ

1. **ネットワークエラー**
   - Genkit サーバーが起動しているか確認
   - CORS設定が正しいか確認
   - `.env.local` の API_BASE_URL が正しいか確認

2. **認証エラー**
   - 開発環境では自動的にモック認証が使用されます
   - 本番環境では正しいFirebase設定が必要

3. **ストリーミングエラー**
   - SSE（Server-Sent Events）の形式を確認
   - Genkit Flow のレスポンス形式を確認

## テスト

```bash
# ユニットテストの実行
npm test

# E2Eテストの実行（Playwright）
npm run test:e2e
```

## ビルドとデプロイ

```bash
# プロダクションビルド
npm run build

# ビルドの確認
npm run start
```

## トラブルシューティング

### よくある問題

1. **"Failed to fetch" エラー**
   - Genkit サーバーが起動しているか確認
   - ポート 9090 が使用されているか確認

2. **CORS エラー**
   - サーバー側の CORS 設定を確認
   - `Access-Control-Allow-Origin` ヘッダーが設定されているか

3. **ストリーミングが動作しない**
   - `Accept: text/event-stream` ヘッダーが送信されているか
   - サーバーが SSE 形式で応答しているか

## ライセンス

[ライセンス情報を記載]