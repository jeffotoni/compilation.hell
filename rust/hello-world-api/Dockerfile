FROM rust:1.46.0 AS builder

WORKDIR /project
COPY Cargo.toml ./
COPY src ./src
RUN rustup target add x86_64-unknown-linux-musl
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-musl

# # Bundle Stage
FROM scratch
COPY --from=builder /project/target/x86_64-unknown-linux-musl/release/hello-world-api .
CMD ["./hello-world-api"]