# Stage 1

FROM node:22-alpine AS build



WORKDIR /srv



COPY package.json package-lock.json ./



RUN npm ci



COPY . .



RUN npm run build



# Stage 2

FROM node:22-alpine



ENV NODE_ENV=production



WORKDIR /srv



COPY package.json package-lock.json ./

RUN npm ci --omit=dev



COPY --from=build /srv/dist ./dist



EXPOSE 3000



HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \

  CMD node -e "fetch('http://127.0.0.1:3000/health', {signal: AbortSignal.timeout(3000)}).then(r => process.exit(r.ok ? 0 : 1)).catch(() => process.exit(1))"



USER node



CMD ["npm", "run", "start"]
