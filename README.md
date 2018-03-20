# PorSimplesSent
## A Portuguese corpus of aligned sentences pairs to investigate sentence readability assessment

## NILC
This corpus was created during my master's degree at ICMC-USP, and made possible thanks to the Interinstitutional Center for Computational Linguistics - NILC (Núcleo Interestitucional de Linguística Computacional), represented by my advisor Dra. Sandra Maria Aluísio and the linguistics specialist Dra. Magali Sanches Duran.

[http://www.nilc.icmc.usp.br/nilc/index.php](http://www.nilc.icmc.usp.br/nilc/index.php)


## License
[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)

## TSV format
All files are in Tab Separated Values (TSV) format, it means that fields are separated by tab (Also knows as ```char(9)``` or ```\t```), and newline (```char(10)``` or ```\n```) for the rows.

## PorSimples

In this folder you'll find the source corpus used to extract the sentence pairs, already exportaded in TSV format:

### porsimples_sentences.tsv
* __production_id__: Each triplet of texts (original, natural, strong) has an unique id, called production_id.
* __level__: ORI (1 - Original), level NAT (2 - Natural) or STR (3 - Strong)
* __text_id__: Unique id for each text.
* __sentence_id__: Unique id for each sentence.
* __paragraph__: Sequential id for the paragraph in text.
* __sentence_text__: The raw text from the sentence.

### porsimples_aligns.tsv
* __production_id__: See porsimples_sentences.tsv.
* __level__: Simplification level ORI->NAT or NAT->STR.
* __text_id_from__: Text id from source side of simplification.
* __sentence_id_from__: sentence id from source side of simplification.
* __text_id_to__: Text id for target side of simplification.
* __sentence_id_to__: Sentence id for target side of simplification.

## PorSimplesSent (pss)

In this folder are the files with aligned pairs from pss0 to pss3, it all have the same layout:
* __production_id__: See porsimples_sentences.tsv.
* __level__: Simplification level ORI->NAT, NAT->STR or ORI->STR.
* __changed__: If the sentence has changes in this simplification level.
* __split__: If the sentence suffers split in this simplification level.
* __sentence_text_from__: The raw text of the source sentence.
* __sentence_text_to__: The raw text of the target sentence.

### pss0 - Split sentences concatenated
Concatenate all resulting split sentences on the right side, may be usefull to study the simplification process.
* pss0_align_concat_ori_nat.tsv
* pss0_align_concat_nat_str.tsv

### pss1 - All splits (1 to n)
Repeats left side sentence to each one resulting split
* pss1_align_all_splits_ori_nat.tsv
* pss1_align_all_splits_nat_str.tsv
* pss1_align_all_splits_ori_str.tsv

### pss2 - Major Length splits (1 to major(n))
Only the sentence with bigger length and most overlap of tokens. Repeats left side sentence when two resulting split sentences has the same size and overlap.
* pss2_align_length_ori_nat.tsv
* pss2_align_length_nat_str.tsv
* pss2_align_length_ori_str.tsv

### pss3 - No split sentences (1 to 1)
Only the sentences that not suffered split.
* pss3_align_no_splits_ori_nat.tsv
* pss3_align_no_splits_nat_str.tsv
* pss3_align_no_splits_ori_str.tsv

## PorSimplesSent - Triplets
In the file triplets_length.tsv, are sentences from the 3 levels, generated from the pss2_length pairs, in the following layout:
* __production_id__: See porsimples_sentences.tsv.
* __level__: Fixed - ORI->NAT->STR.
* __changed_ori_nat__: If the sentence has changes from the original to the natural level.
* __changed_nat_str__: If the sentence has changes from the natural to the strong level.
* __original_text__: The raw text of the original sentence.
* __natural_text__: The raw text of the natural sentence.
* __strong_text__: The raw text of the strong sentence.

## Statistics
```
Total sentences Original: 2915
      Zero Hora: 2071
      Caderno Ciencia FSP: 844
Total sentences Natural: 4072
Total sentences Strong: 4977
Total sentences ALL: 11964

Total sentences NO SIMPLIFICATION Original->Natural: 542
Total sentences NO SIMPLIFICATION Natural->Strong: 2575

Total sentences SPLIT Original->Natural: 829
Total sentences SPLIT Natural->Strong: 720

Total sentences Natural from split: 1992
Total sentences Strong from split: 1623

Total sentences SIMPLIFIED (no split) Original->Natural: 1543
Total sentences SIMPLIFIED (no split) Natural->Strong: 779

Total pairs simplified Original->Natural: 2372
Total pairs simplified Natural->Strong: 1499
Total pairs simplified Original->Strong: 1139
Total all pairs simplified: 5010

Total triplets NO SIMPLIFICATION 3 Levels: 372
Total triplets Simplified Only Original->Natural: 1291
Total triplets Simplified Only Natural->Strong: 178
Total triplets Simplified 3 Levels: 1139
Total triplets: 2980

Mean token size of sentences - simplified (no split) - Ori->Nat: 20
Min token size of sentences - simplified (no split) - Ori->Nat: 3
Max token size tokens of sentences - simplified (no split) - Ori->Nat: 69

Mean token size of sentences - simplified (with split) - Ori->Nat: 33
Min token size of sentences - simplified (with split) - Ori->Nat: 6
Max token size tokens of sentences - simplified (with split) - Ori->Nat: 54

Mean token size of sentences - simplified (no split) - Nat->Str: 22
Min token size of sentences - simplified (no split) - Nat->Str: 4
Max token size tokens of sentences - simplified (no split) - Nat->Str: 57

Mean token size of sentences - simplified (with split) - Nat->Str: 24
Min token size of sentences - simplified (with split) - Nat->Str: 5
Max token size tokens of sentences - simplified (with split) - Nat->Str: 49

Mean tokens size diff of sentences - Originals vs simplified (no split) - Ori->Nat: 6
Min tokens size diff of sentences - Originals vs simplified (no split) - Ori->Nat: 1
Max tokens size diff of sentences - Originals vs simplified (no split) - Ori->Nat: 27

Mean tokens size diff of sentences - Originals vs simplified (with split) - Ori->Nat: 9
Min tokens size diff of sentences - Originals vs simplified (with split) - Ori->Nat: 1
Max tokens size diff of sentences - Originals vs simplified (with split) - Ori->Nat: 64

Total PSS1 Original->Natural: 3535
Total PSS1 Natural->Strong: 4977
Total PSS1 Original->Strong: 2104
Total geral PSS1: 10616

Total PSS2 Original->Natural: 2372
Total PSS2 Natural->Strong: 1499
Total PSS2 Original->Strong: 1093
Total geral PSS2: 4964

Total PSS3 Original->Natural: 1543
Total PSS3 Natural->Strong: 779
Total PSS3 Original->Strong: 272
Total geral PSS3: 2594
```
