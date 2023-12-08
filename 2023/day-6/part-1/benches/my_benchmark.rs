use criterion::{criterion_group, criterion_main, Criterion};
use part_1::{process, process_v2};

fn my_benchmark_function(c: &mut Criterion) {
    let input = include_str!("../src/input.txt");

    c.bench_function("process", |b| {
        b.iter(|| {
            process(input)
        });
    });

    c.bench_function("processV2", |b| {
        b.iter(|| {
            process_v2(input)
        });
    });
}

criterion_group!(benches, my_benchmark_function);
criterion_main!(benches);
