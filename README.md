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

